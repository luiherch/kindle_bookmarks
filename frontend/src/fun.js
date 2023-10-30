// document.addEventListener('DOMContentLoaded', function() {
//     const themeButton = document.getElementById('themeButton');
//     const htmlTag = document.querySelector('html');

//     themeButton.addEventListener('click', function() {
//         const currentTheme = htmlTag.getAttribute('data-theme');

//         if (currentTheme === 'light') {
//             htmlTag.setAttribute('data-theme', 'dark');
//         } else {
//             htmlTag.setAttribute('data-theme', 'light');
//         }
//     });
// });

import { Sync } from "../wailsjs/go/main/Library";

import { UpdateConfig } from "../wailsjs/go/main/DbConfig";

window.sync = function () {
  Sync().then((result) => {
    console.log("Synced!")
  });
}

window.update_db = function () {
  let val = document.getElementById("dbid").value
  UpdateConfig("databaseId",val).then((result) => {
    console.log("DbId Updated")
  });
}

window.update_secret = function () {
  let val = document.getElementById("apisecret").value
  UpdateConfig("apiSecret",val).then((result) => {
    console.log("ApiSecret Updated")
  });
}