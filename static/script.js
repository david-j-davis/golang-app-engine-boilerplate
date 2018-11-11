"use strict";

const fetchMessage = () => {
    fetch('/hello', { credentials: 'include' })
    .catch(err => console.log(err))
    .then(resp => resp.text())
    .then(data => {
        document.getElementById('message').innerHTML = data
    })
}