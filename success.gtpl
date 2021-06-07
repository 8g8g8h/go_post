<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>success</title>

</head>
<body>
    <p>登録完了しました</p>
    <form action="/">
        <table>
            <tr>
                <td>単語名：</td>
                <td>{{.name}}</td>
            </tr>
            <tr>
                <td>意味：</td>
                <td>{{.mean}}</td>
            </tr>
        </table>
        <p><button type="submit" formaction="http://localhost:8080/add">追加画面に戻る</button></p>
    </form>
</body>
</html>