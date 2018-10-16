<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>简易聊天室</title>
    <style>
        #content::-webkit-scrollbar{ display:none; }
        #content::-webkit-scrollbar-track{ display:none; }
        #content::-webkit-scrollbar-thumb{ display:none; }
        #content{  overflow:scroll;border: 1px solid;border-radius:9px;width: 700px;height: 500px;  }
        #msg{  height: 50px;width: 500px;border:1px solid;padding-left: 10px;border-radius: 5px;  }
        .red{color: red;}
        .green{color: green;}
        .blue{color: blue;}
        .div1{background-color: rgba(150,15,150,0.1); border: 1px solid; border-radius:9px;display: inline-block;padding:5px 8px;}
        .img{width:30px;height:30px;border: 1px solid;border-radius:100% ;}
    </style>
</head>
<script type="text/javascript" src="/static/js/jquery.min.js"></script>
<script type="text/javascript">
    if (typeof console == "undefined") {    this.console = { log: function (msg) {  } };}
    // 如果浏览器不支持websocket，会使用这个flash自动模拟websocket协议，此过程对开发者透明
    WEB_SOCKET_SWF_LOCATION = "/talk/swf/WebSocketMain.swf";
    // 开启flash的websocket debug
    WEB_SOCKET_DEBUG = true;
    $(document).keyup(function(event){
        if(event.keyCode ==13){
            // var msg = document.getElementById("msg").value;
            //alert(msg);
            //console.log(msg.replace(/\n/g,'<br />'));
           $("#send").trigger("click");
           }
        });
    var ws, name,img, client_list={};

    // 连接服务端
    function connect() {
        // 创建websocket
        ws = new WebSocket('ws://' + window.location.host + '/ws');
        // 当socket连接打开时，输入用户名
        ws.onopen = onopen;
        // 当有消息时根据消息类型显示不同信息
        ws.onmessage = onmessage;
        ws.onclose = function() {
            console.log("连接关闭，定时重连");
            connect();
        };
        ws.onerror = function() {
            console.log("出现错误");
        };
    }

    // 连接建立时发送登录信息
    function onopen()
    {
        if(!name || !img)
        {
            show_prompt();
        }
        // 登录
        var login_data = '{"type":"login","client_name":"'+name.replace(/"/g, '\\"')+'","client_img":"'+img+'"}';
        console.log("websocket握手成功，发送登录数据:"+login_data);
        ws.send(login_data);
    }
    // 输入姓名
    function show_prompt(){
        name = prompt('输入你的名字：', '');
        if(!name || name=='null'){
            name = '游客';
        }
        if(!img || img=='null'){
            img='/static/img/tx'+Math.ceil(Math.random()*3)+'.bmp';
        }
    }
    function onSubmit(){
            var msg = document.getElementById("msg").value;
                      document.getElementById("msg").value="";
            var login_data = '{"type":"say","client_name":"'+name.replace(/"/g, '\\"')+'","content":"'+msg.replace(/\n/g,'')+'","client_img":"'+img+'"}';
            //console.log("消息发送成功，发送数据:"+login_data);
            ws.send(login_data);
    }
    // 服务端发来消息时
    function onmessage(e)
    {
        var data=JSON.parse(e.data);

        switch(data.type)
        {
            case 'login':
                document.getElementById("content").innerHTML += "<h5 class='green'> {" +data.client_name+"}&nbsp;"+ data.content + "</h5>";
                break;
            case 'quit':
                document.getElementById("content").innerHTML += "<h5 class='red'> {" + data.client_name+"}&nbsp;"+ data.content+ "</h5>";
                break;
            case 'say':
                document.getElementById("content").innerHTML += "<table><tr><td rowspan='2'><img class='img' src='"+ data.client_img+ "'></td><td>"+data.client_name+"：</td></tr><tr><td><div class='div1'>" + data.content + "</div></td></tr></table>";
                break;
        }

         var scrollDiv=document.getElementById("content");
             scrollDiv.scrollTop = scrollDiv.scrollHeight;
    }

</script>
<body onload="connect();">
<div style="margin-left: 25%;">
<h1>极简聊天室</h1>（基于workerman）
<div id="content"></div>
<br>
<a href="javascript:void(0);">点击发送表情</a>
<br><br>
    <textarea id="msg"></textarea>
    <button type="button" id="send" onclick="onSubmit()">发送</button>（回车或单击）
</div>
</body>
</html>
