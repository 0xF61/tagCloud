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
	.font-11{font-size: 20px }
	.font-12{font-size: 24px }
	.font-13{font-size: 28px }
	.font-14{font-size: 32px }
	.font-15{font-size: 36px }
	.font-16{font-size: 40px }
	.font-17{font-size: 44px }
	.font-18{font-size: 48px }
	.font-19{font-size: 52px }
	.font-20{font-size: 56px }

	div{
		width: 800px;
		height: 500px;
		position: absolute;
		top:0;
		bottom: 0;
		left: 0;
		right: 0;
		margin: auto;
	}
	</style>
</head>
<body>
<center>
<a href="/">Index</a> |
<a href="/tag">Tags</a>
</center>
<br>
<div class=tags>
{{range .Tags}}
<span class="font-{{.Size}}">{{.Title}}</span>({{.Freq}})  {{end}}
</div>
</body>
</html>
