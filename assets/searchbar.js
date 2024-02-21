// Récupérer la valeur de 'q'
const url = window.location.href;
const urlParams = new URLSearchParams(window.location.search);
const queryValue = urlParams.get('q');
var firstSearch = true

// Récupération de tt les éléments utiles
var searchtext = document.getElementById('artist-search');
var range1 = document.getElementById('range');
var range2 = document.getElementById('range2');
range1.value = "1958"
range2.value = "2015"
var tri = document.getElementById("tri");
var triNB = document.getElementById("triNB");
var texteBar = document.getElementsByClassName('queryValue');
var texteDate1 = document.getElementsByClassName('date1');
var texteDate2 = document.getElementsByClassName('date2');
texteBar[0].textContent = "Tous les artistes";
texteDate1[0].textContent = range1.value;
texteDate2[0].textContent = range2.value;

//On trie par ordre alphabétique
let artists = Array.from(document.querySelectorAll('.cards'));
const artistList = document.querySelector('.cardsList');
artists.sort((a, b) => {
    return a.querySelector('.name').textContent.localeCompare(b.querySelector('.name').textContent);
});
artistList.innerHTML = ''; 

artists.forEach(function(artist) {
    artistList.appendChild(artist);
});


//Si la recherche provient de la page d'accueil
if (queryValue != '' && firstSearch) {
    firstSearch = false
    searchtext.value = queryValue
    texteBar[0].textContent = searchtext.value;
    Filtered();
}

searchtext.addEventListener('input',function(){
    texteBar[0].textContent = searchtext.value;
    if (texteBar[0].textContent == "") {
        texteBar[0].textContent = "Tous les artistes";
    }
    Filtered();
});

range1.addEventListener('input',function(){
    texteDate1[0].textContent = range1.value;
    Filtered();
});

range2.addEventListener('input',function(){
    texteDate2[0].textContent = range2.value;
    Filtered();
});

tri.addEventListener('change',function(){
    const triValue = tri.value;
    // Trier les artists en fonction de l'option de tri sélectionnée
    if (triValue == 'true') {
        artists.sort((a, b) => {
            return a.querySelector('.name').textContent.localeCompare(b.querySelector('.name').textContent);
        });
    } else if (triValue == 'false') {
        artists.sort((a, b) => {
            return b.querySelector('.name').textContent.localeCompare(a.querySelector('.name').textContent);
        });
    }
    artistList.innerHTML = ''; 

    artists.forEach(function(artist) {
        artistList.appendChild(artist);
    });
    Filtered();
});

triNB.addEventListener('change',function(){
    Filtered();
});

function Filtered() { // Fonction qui trie/filtre les artistes en fonction des values des input
    var text = searchtext.value.toLowerCase();
    for (var i = 0; i < artists.length; i++) {
        var name = artists[i].querySelector('.name').textContent.toLowerCase();
        var creationDate = artists[i].querySelector('.date').value;
        var members = artists[i].querySelector('.member').value;
        var NBmembers = artists[i].querySelector('.NBmembers').value;
        if ((name.includes(text)||creationDate.includes(text)||members.toLowerCase().includes(text))&&(texteDate1[0].textContent <= creationDate && texteDate2[0].textContent >= creationDate)) {
            if (triNB.value != "0") {
                if (triNB.value == NBmembers) {
                    artists[i].style.display = 'block';
                } else {
                    artists[i].style.display = 'none';
                }
            } else {
                artists[i].style.display = 'block';
            }
        } else {
            artists[i].style.display = 'none';
        }
    }
}
