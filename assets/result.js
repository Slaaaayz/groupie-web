function getQueryParam(name) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(name);
}

const result = getQueryParam('result');
const nameArtist = decodeURIComponent(getQueryParam('name'));
const idArtist = decodeURIComponent(getQueryParam('idArtist'));
const image = decodeURIComponent(getQueryParam('image'));
alert(image)
document.getElementById('gameResult').textContent = result === 'win' ? 'Bravo!' : 'Dommage...';
document.getElementById('artistName').textContent = nameArtist;
// Correction pour éviter la duplication et utiliser correctement l'ID
document.getElementById('artistIdSpan').textContent = idArtist;
document.getElementById('artistLink').href = `/artist/${idArtist}`;

const imgElement = document.getElementById('artistImage');
if (image && image.trim() !== '') {
    imgElement.src = image.trim;
    imgElement.alt = `Image de ${nameArtist}`;
    imgElement.onerror = function() {
        console.error("Erreur lors du chargement de l'image :", image);
        this.style.display = 'none'; // Cache l'image si elle ne peut pas être chargée
    };
} else {
    console.log("Aucune URL d'image fournie.");
    imgElement.style.display = 'none';
}


