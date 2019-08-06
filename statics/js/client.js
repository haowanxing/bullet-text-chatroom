class Client {
    constructor(handler){
        this.ws = null;
        this.handler = handler;
        this.name = "匿名" + Date.now().toString();
        this.room = "default";
        this.onMessageList = [];
        this.onCloseList = [];
        this.onOpenList = [];
        this.onErrorList = [];
    }
    //发送消息
    sendText(text){
        let msg = {
            "event":"text",
            "data":text
        };
        this.ws.send(JSON.stringify(msg));
    };
    changeName(name){
        let msg = {
            "event":"name",
            "data":name
        };
        this.ws.send(JSON.stringify(msg));
        this.name = name
    };
    changeRoom(room){
        let msg = {
            "event":"room",
            "data":room
        };
        this.ws.send(JSON.stringify(msg));
    };
    connect() {
        let _this = this;
        this.ws = new WebSocket("ws://" + window.location.host + "/ws");
        this.ws.onmessage = function (e) {
            let data = eval('(' + e.data + ')');
            if (typeof _this.handler == "function") {
                _this.handler(data);
            }
            for (let i = 0; i < _this.onMessageList.length; i++) {
                let fn = _this.onMessageList[i];
                if (typeof fn == "function") {
                    fn(data);
                }
            }
        };
        this.ws.onclose = function(e){
            console.log("ws已断开！");
            _this.reconnect();
            for (let i = 0; i < _this.onCloseList.length; i++) {
                let fn = _this.onCloseList[i];
                if (typeof fn == "function") {
                    fn(e);
                }
            }
        };
        this.ws.onopen = function (e) {
            console.log("ws已连接！");
            for (let i = 0; i < _this.onOpenList.length; i++) {
                let fn = _this.onOpenList[i];
                if (typeof fn == "function") {
                    fn(e);
                }
            }
        };
        this.ws.onerror = function (e) {
            console.log("发生错误："+e);
            for (let i = 0; i < _this.onErrorList.length; i++) {
                let fn = _this.onErrorList[i];
                if (typeof fn == "function") {
                    fn(e);
                }
            }
        };
    }
    reconnect(){
        let _this = this;
        setTimeout(function(){
            _this.connect();
        },1000);
    }

    wsOnMessage(fn){
        this.onMessageList.push(fn)
    }
    wsOnClose(fn){
        this.onCloseList.push(fn)
    }
    wsOnOpen(fn){
        this.onOpenList.push(fn)
    }
    wsOnError(fn){
        this.onErrorList.push(fn)
    }
}