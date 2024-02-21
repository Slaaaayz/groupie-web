document.addEventListener('DOMContentLoaded', function() {
    let currentHint = 0;
    const hints = document.querySelectorAll('.textIndice p, .textIndice img');
    const chrono = document.querySelector('.chrono');
    // Assurez-vous que ces éléments existent dans votre HTML avec les attributs data-value corrects
    const artistName = document.getElementById('artistName').getAttribute('data-value').trim();
    const artistId = document.getElementById('idArtist').getAttribute('data-value').trim(); 
    const artistImage = document.getElementById('artistImage').getAttribute('data-value').trim();
    let attempts = 3;
    let timeLeft = 10;

    hints.forEach((hint, index) => {
        if (index > 0) hint.style.display = 'none';
    });

    const updateChrono = () => {
        if (timeLeft > 0) {
            timeLeft -= 1;
            chrono.textContent = `00:${timeLeft < 10 ? '0' : ''}${timeLeft}`;
        } else {
            nextHint();
        }
    };

    let timer = setInterval(updateChrono, 1000);

    function nextHint() {
        if (currentHint < hints.length - 1) {
            hints[currentHint].style.display = 'none';
            currentHint += 1;
            hints[currentHint].style.display = 'block';
            timeLeft = 10;
        } else {
            redirectToResultPage('lose');
        }
    }

    function redirectToResultPage(result) {
        const artistNameEncoded = encodeURIComponent(artistName);
        const artistImageEncoded = encodeURIComponent(artistImage);
        const artistIdEncoded = encodeURIComponent(artistId);
        const resultPageURL = `/result?result=${result}&name=${artistNameEncoded}&image=${artistImageEncoded}&idArtist=${artistIdEncoded}`;
        window.location.href = resultPageURL;
    }

    document.querySelector('button').addEventListener('click', function() {
        const userInput = document.getElementById('artist-search').value;
        if (userInput.toLowerCase() === artistName.toLowerCase()) {
            redirectToResultPage('win');
        } else {
            attempts -= 1;
            if (attempts === 0) {
                redirectToResultPage('lose');
            } else {
                alert(`Faux! Il vous reste ${attempts} tentatives.`);
                nextHint();
            }
        }
    });
});
