
document.addEventListener('DOMContentLoaded', function() {
    const articleElement = document.querySelector('details');
    if (!articleElement) {
        console.log('No books found');
        const modal = document.getElementById("highlights-modal");
        openModal(modal)
    }
  });


