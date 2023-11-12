const articleElement = document.querySelector('article');

document.addEventListener('DOMContentLoaded', function() {
    if (!articleElement) {
        console.log('No books found');
        const modal = document.getElementById("highlights-modal");
        openModal(modal)
    }
  });


