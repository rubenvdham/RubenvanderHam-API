'''
Created on 16 nov. 2017

@author: Ruben
'''

from bottle import run,route, error ,request,response, get, put, delete
import time
from json import dumps
myEpochAge = 894697740



@get('/myip')
def returnIP():
    response.content_type = 'text/html; charset=UTF-8'
    return "<p>"+request.get('REMOTE_ADDR')+"</p>"
    
@get("/")
def hello():
    return "Welcome to the API, For documentation visit '/documentation'"

@get("/age")
def getAge():
    currentTime = time.time()
    differenceDays=((currentTime-myEpochAge)/3600)/24
    JSON = [{"years":round(differenceDays/365,2), "days":int(differenceDays)}]
    response.content_type = 'application/json'
    return dumps(JSON)

@get("/documentation")
def dumpDocumentation():
    response.content_type = 'application/json'
    JSON = [
        {"location":"/age","methods":"GET","returnType":"JSON","description":"Returns my age with keys 'years' and 'days' and their corresponding values"},
        {"location":"/myip","methods":"GET","returnType":"HTML","description":"Returns the IP address of the client"},
        {"location":"/documentation", "methods":"GET", "returnType":"JSON", "description":"Returns the documentation of this API in JSON"}]
    return dumps(JSON)


@error(404)
def error404(error):
    response.content_type = 'text/html; charset=UTF-8'
    return '<p>Nothing here, sorry<br> Try \'\documentation\' for more info</p>'

@error(500)
def error500(error):
    response.content_type = 'text/html; charset=UTF-8'
    return 'Shitt, something went wrong....'

if __name__ == '__main__':
    run(host='0.0.0.0', port=8001,reloader=True)