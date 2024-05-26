<template>
  <router-view />
</template>
<style>
@font-face {
  font-family: "Bahij TheSansArabic";
  font-weight: 500;
  src: url("./assets/BahijTheSansArabic-Bold.ttf");
}
@font-face {
  font-family: "Bahij TheSansArabic";
  font-weight: 800;
  src: url("./assets/Bahij_TheSansArabic-ExtraBold.ttf");
}
body {
  font-family: "Bahij TheSansArabic";
  font-weight: 500;
  margin: 0;
  padding: 0;
  min-height: 100vh;
  width: 100%;
  box-sizing: border-box;
  scroll-behavior: auto;
  touch-action: manipulation;
  border: #6835f8 10px solid;
  background: linear-gradient(
    60deg,
    rgb(128 118 235/0.4),
    rgb(175 170 227/0.2),
    rgb(175 170 227/0.1),
    rgb(255, 255, 255)
  );
}
button {
  background: none;
  color: inherit;
  border: none;
  padding: 0;
  font: inherit;
  cursor: pointer;
  outline: inherit;
  transition: 0.5s;
}
button:hover {
  box-shadow: 0 0 15px rgba(74, 20, 222, 0.4);
}
</style>
