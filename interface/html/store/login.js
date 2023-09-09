let _UserToken = ""

$(document).ready(function() {
    IsLoginHtml()
});

function SetLoginHtml(data){
    localforage.setItem('user_token', JSON.stringify(data))
}

function IsLoginHtml() {
    // 回调版本：从仓库中获取 key 对应的值并将结果提供给回调函数。如果 key 不存在，getItem() 将返回 null。
    localforage.getItem('user_token', function(err, value) {
        // 当离线仓库中的值被载入时，此处代码运行
        let res = JSON.parse(value)
        let now = getCurrentFormattedDateTime()
        console.log(now , res.exp)
        if (res.exp <= now ){
            localforage.clear()
            _UserToken = ""
            HeaderLoginPage()
            return
        }
        _UserToken = res.token
        console.log(_UserToken);
        HeaderLoginPage()
    });
}
function getCurrentFormattedDateTime() {
    var now = new Date();
    var year = now.getFullYear();  // 年份
    var month = padZero(now.getMonth() + 1);  // 月份（注意月份从0开始计数）
    var day = padZero(now.getDate());  // 日期
    var hours = padZero(now.getHours());  // 小时
    var minutes = padZero(now.getMinutes());  // 分钟
    var seconds = padZero(now.getSeconds());  // 秒钟

    return year + '-' + month + '-' + day + ' ' + hours + ':' + minutes + ':' + seconds;
}

function padZero(number) {
    return number < 10 ? '0' + number : number;
}

function HeaderLoginPage(){
    console.log("HeaderLoginPage", _UserToken)
    if (_UserToken === "") {
        $("#isLogin").show()
        $("#isLoginTrue").hide()
        return
    }

    // 获取用户信息
    $.post("/login/check", { "token":_UserToken},
        function(res){
            console.log(res)
            if (res.code !== 0 ){
                layer.msg(res.message)
                if (res.code  === -126) {
                    localforage.clear()
                }
                $("#isLogin").show()
                $("#isLoginTrue").hide()
                return
            }

            // 设置头部信息
            $("#isLogin").hide()
            $("#userNmaeSee").html(res.data.name+"("+res.data.ip+")")
            $("#isLoginTrue").show()
        });
}

function LoginOut(){
    localforage.clear()
    _UserToken = ""
    $("#isLogin").show()
    $("#userNmaeSee").html("")
    $("#isLoginTrue").hide()
}