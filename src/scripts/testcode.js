const host = "http://localhost:8080/";
function createHost(path) {
    return host + path;
}

//аутентификация
// var s;
// fetch('http://localhost:8080/login', {
//     method: 'POST',
//     body: JSON.stringify({
// 	    username: 'testuser',
// 	    password: 'test'
//     }),
//     headers: {
//         "Content-Type": "application/json",
//     }
// })
// .then((res) => {
//     console.log(res.headers.get('set-cookie'));
//     s = res.headers.get('set-cookie');
//     if(res.ok) {
//         return res.text();
//     }
//     return Promise.reject(res.status);
// })
// .then((res) => {
//     console.log(res);
// })
// .catch(err => console.log(`Ошибка: ${err}`));

fetch(createHost(`api/user/tasks/categories`), {
    method: 'GET',
    headers: {
        "Content-Type": "application/json",
    },
})
.then((res) => {
  if(res.ok) {
    return res.json()
  }
  return Promise.reject(res.status);
})
.then((res) => {
  console.log(res)
})
.catch(err => console.log(`Ошибка: ${err}`));
