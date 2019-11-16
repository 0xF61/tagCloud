<!doctype html>
<html lang="en-US">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>{{.PageTitle}}</title>
  <link rel="shortcut icon" href="favicon.ico">
  <link rel="icon" href="favicon.ico">
  <link rel="stylesheet" type="text/css" href="styles.css">
  <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
	<style>
	body{
		background-color:black;
		color:#ccc;
	}
	a{
		color:yellow;
	}
	</style>
</head>

<body>
<center>
<a href="/">Index</a> |
<a href="/tag">Tags</a>
<br>
<h1>{{.Context}}</h1>

<form action="/tag", method="POST">
  URL:<br>
  <input type="text" name="url" value="https://tr.lipsum.com/">
  <br>
  <input type="text" name="minFrequency" value="3">
  <br><br>
  <input type="submit" value="Submit">
</form>

</center>
</body>
</html>
