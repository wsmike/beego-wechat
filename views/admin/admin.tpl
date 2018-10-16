<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  
</head>

<body>
    <h1> 登陆页面 {{ .id }}</h1>
    <form action="/allPost" method="post">
	    <input type="hidden" name="_xsrf" value="{{ .xsrf }}">
	
 		<input type="submit" value="提交" />
    </form>
</body>
</html>
