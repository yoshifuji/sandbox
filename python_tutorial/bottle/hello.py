# -*- coding:utf-8 -*-

from bottle import route, run


@route('/hello/')
@route('/hello/<user>')
def hello(user="taro"):
    return "Hello {user}".format(user=user)


@route('/date/<month:re:[a-z]+>/<day:int>/<path:path>')
def date(month, day, path):
    return "{month}/{day} {path}".format(month=month, day=day, path=path)

run(host='localhost', port=8080, debug=True)
