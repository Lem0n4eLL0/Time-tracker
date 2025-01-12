// //регестрация
fetch('http://localhost:8080/registration/', {
    method: 'POST',
    body: JSON.stringify({
        email: 'vladislavch31e@bk.ru',
	    username: 'vladik3',
	    password: 'qwerty3'
    }),
    headers: {
        "Content-Type": "application/json",
    }
})
.then((res) => {
    if(res.ok) {
        return res.text();
    }
    return Promise.reject(res.status);
})
.then((res) => {
    console.log(res);
})
.catch(err => console.log(`Ошибка: ${err}`));

//аутентификация
// fetch('http://localhost:8080/login/', {
//     method: 'POST',
//     body: JSON.stringify({
// 	    username: 'vladik2',
// 	    password: 'qwerty2'
//     }),
//     headers: {
//         "Content-Type": "application/json",
//     }
// })
// .then((res) => {
//     console.log(res) //сохранять токен
//     if(res.ok) {
//         return res.json();
//     }
//     return Promise.reject(res.status);
// })
// .then((res) => {
//     console.log(res);
// })
// .catch(err => console.log(`Ошибка: ${err}`));
// Проверка токена
// fetch('http://localhost:8080/profile/', {
//     method: 'POST',
//     body: JSON.stringify({
// 	    username: 'vladik2',
// 	    password: 'qwerty2'
//     }),
//     headers: {
//         "Content-Type": "application/json",
//         "authorization": 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ2bGFkaWsiLCJleHAiOjE3MzY2MDQ1NTAsImFkbWluIjpmYWxzZSwiZGF0ZSI6bnVsbH0.21u3WMcS8Bx9a_NDi46-4EJRWbBOC95mexhFk0QSEYo'
//     }
// })
// .then((res) => {
//     console.log(res)
//     if(res.ok) {
//         return res.text();
//     }
//     return Promise.reject(res.status);
// })
// .then((res) => {
//     console.log(res);
// })
// .catch(err => console.log(`Ошибка: ${err}`));