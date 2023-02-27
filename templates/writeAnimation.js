const txtAnim = document.querySelector('h1');

const secondTxtAnim = document.querySelector('h2');

new Typewriter(txtAnim, {
     deleteSpeed: 15,
})


.changeDelay(200)
.typeString('<span style="color: #ffffff">Bienvenue sur notre Groupie Tracker </span>')
.pauseFor(300)
.start()

new Typewriter(secondTxtAnim, {
     deleteSpeed: 15,
})
.pauseFor(9000)
.typeString('<span style="color: #ffffff">Réalisé par :</span>')
.pauseFor(750)
.typeString('<span style="color: #27ae60"> Tayvadi</span')
.deleteChars(8)
.pauseFor(750)
.typeString('<span style="color: #27ae60"> Tom </span>')
.pauseFor(750)
.deleteChars(4)
.typeString('<span style="color: #27ae60"> Charles </span>')
.deleteChars(24)
.start()
