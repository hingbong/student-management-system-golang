const sexArr = ['男', '女'];

const confirm_delete = url => {
    if (confirm("确定要删除么？")) {
        fetch(url, {
            method: 'delete'
        }).then(response => response.json().then(json => {
            if (json.code === 1) {
                window.location.reload();
            } else {
                alert(json.message);
            }
        }))
    }
};

const search = form => {
    const body = new URLSearchParams();
    Array.from(form).forEach(e => body.append(e.name, e.value));
    fetch(`${form.action}?${body}`, {
        method: 'get'
    }).then(
        response => response.json().then(json => putData(json)))
};

const send_form = form => {
    const body = new URLSearchParams();
    Array.from(form).forEach(e => body.append(e.name, e.value));
    fetch(form.action, {
        method: 'post',
        body: body
    }).then(response => response.json().then(json => {
        const notice = document.querySelector('#form_notice');
        if (json.code === 1) {
            notice ? notice.innerText = '成功' : '';
            window.location.replace('index.html');
        } else if (notice) {
            notice.innerText = json.message;
        }
    }));
};

const modify = form => {
    const body = new URLSearchParams();
    Array.from(form).forEach(e => body.append(e.name, e.value));
    fetch(form.action, {
        method: 'put',
        body: body
    }).then(response => response.json().then(json => {
        const notice = document.querySelector('#form_notice');
        if (json.code === 1) {
            notice ? notice.innerText = '成功' : '';
            window.location.replace('index.html');
        } else if (notice) {
            notice.innerText = json.message;
        }
    }));
};

const getQueryStringParameters = () => {
    let query = window.location.search.substring(1);
    return (/^[?#]/.test(query) ? query.slice(1) : query)
        .split('&')
        .reduce((params, param) => {
            let [key, value] = param.split('=');
            params[key] = value ? decodeURIComponent(value.replace(/\+/g, ' '))
                : '';
            return params;
        }, {});
};

const parseAge = identityCard => {
    const len = (identityCard + "").length;
    let strBirthday = "";
    //处理18位的身份证号码从号码中得到生日和性别代码
    if (len === 18) {
        strBirthday = identityCard.substr(6, 4) + "/" + identityCard.substr(10,
            2) + "/" + identityCard.substr(12, 2);
    }
    if (len === 15) {
        let birthdayValue;
        birthdayValue = identityCard.charAt(6) + identityCard.charAt(7);
        if (parseInt(birthdayValue) < 10) {
            strBirthday = "20" + identityCard.substr(6, 2) + "/"
                + identityCard.substr(8, 2) + "/" + identityCard.substr(10, 2);
        } else {
            strBirthday = "19" + identityCard.substr(6, 2) + "/"
                + identityCard.substr(8, 2) + "/" + identityCard.substr(10, 2);
        }

    }
    //时间字符串里，必须是“/”
    const birthDate = new Date(strBirthday);
    const nowDateTime = new Date();
    let age = nowDateTime.getFullYear() - birthDate.getFullYear();
    //再考虑月、天的因素;.getMonth()获取的是从0开始的，这里进行比较，不需要加1
    if (nowDateTime.getMonth() < birthDate.getMonth() || (nowDateTime.getMonth()
        === birthDate.getMonth() && nowDateTime.getDate()
        < birthDate.getDate())) {
        age--;
    }
    return age;
};