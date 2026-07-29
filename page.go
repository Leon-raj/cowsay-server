package main

const page = `
<!DOCTYPE html>
<html>
<head>
<title>Cowsay</title>
<style>
body{
	font-family: monospace;
	max-width:700px;
	margin:auto;
	padding:40px;
}
textarea{
	width:100%;
	height:80px;
	font-size:16px;
}
pre{
	background:#222;
	color:#0f0;
	padding:20px;
	border-radius:8px;
	overflow:auto;
}
button{
	padding:10px 20px;
	margin-top:10px;
}
</style>
</head>

<body>

<h1>Hello Hopeful Elites!</h1>

<form>
<textarea name="text">{{.Input}}</textarea>

<br>

<button>Speak!</button>
</form>

<pre>{{.Cow}}</pre>

</body>
</html>
`
