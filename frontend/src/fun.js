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
  document.getElementById('sync').disabled = true;
  document.getElementById('sync').setAttribute("aria-busy", "true")
  Sync().then((result) => {
    console.log("Synced!")
    const modal = document.getElementById("confirmation-modal");
    document.getElementById('sync').disabled = false;
  document.getElementById('sync').setAttribute("aria-busy", "false")
    openModal(modal)
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

// Open modal
const openModal = (modal) => {
  if (isScrollbarVisible()) {
    document.documentElement.style.setProperty("--scrollbar-width", `${getScrollbarWidth()}px`);
  }
  document.documentElement.classList.add(isOpenClass, openingClass);
  setTimeout(() => {
    visibleModal = modal;
    document.documentElement.classList.remove(openingClass);
  }, animationDuration);
  modal.setAttribute("open", true);
  console.log("Modal open")
};

const getScrollbarWidth = () => {
  // Creating invisible container
  const outer = document.createElement("div");
  outer.style.visibility = "hidden";
  outer.style.overflow = "scroll"; // forcing scrollbar to appear
  outer.style.msOverflowStyle = "scrollbar"; // needed for WinJS apps
  document.body.appendChild(outer);

  // Creating inner element and placing it in the container
  const inner = document.createElement("div");
  outer.appendChild(inner);

  // Calculating difference between container's full width and the child width
  const scrollbarWidth = outer.offsetWidth - inner.offsetWidth;

  // Removing temporary elements from the DOM
  outer.parentNode.removeChild(outer);

  return scrollbarWidth;
};

// Is scrollbar visible
const isScrollbarVisible = () => {
  return document.body.scrollHeight > screen.height;
};