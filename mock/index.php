<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Ultimate Weirdness</title>
  <style>
    /* Global reset + weird font */
    * { margin:0; padding:0; box-sizing:border-box; }
    body {
      font-family: 'Comic Sans MS', cursive, sans-serif;
      background: linear-gradient(45deg, #ff00cc, #3333ff);
      color: #ffee00;
      overflow-x: hidden;
    }

    /* Rotating header */
    h1 {
      font-size: 4rem;
      text-align: center;
      animation: spin 10s linear infinite;
      mix-blend-mode: difference;
    }
    @keyframes spin {
      from { transform: rotate(0deg); }
      to   { transform: rotate(360deg); }
    }

    /* Card with skew and blur */
    .card {
      width: 300px;
      margin: 2rem auto;
      padding: 1rem;
      background: rgba(0,0,0,0.5);
      backdrop-filter: blur(5px);
      clip-path: polygon(10% 0%, 100% 0%, 90% 100%, 0% 100%);
      transform: skew(-10deg);
      box-shadow: 0 0 20px #00ff99;
      transition: transform 0.3s ease;
    }
    .card:hover {
      transform: skew(0deg) scale(1.05);
    }
    .card p {
      font-size: 0.9rem;
      line-height: 1.4;
      color: #ddffee;
    }

    /* Image gallery with hue-rotate */
    .gallery {
      display: flex;
      gap: 1rem;
      padding: 1rem;
      justify-content: center;
    }
    .gallery img {
      width: 200px;
      height: 150px;
      object-fit: cover;
      filter: hue-rotate(90deg) saturate(2);
      transition: filter 0.5s ease;
    }
    .gallery img:hover {
      filter: none;
    }

    /* Floating footer text */
    footer {
      position: fixed;
      bottom: 0; width: 100%;
      text-align: center;
      background: rgba(0,0,0,0.7);
      color: #ff6699;
      font-size: 1.2rem;
      animation: float 5s ease-in-out infinite;
    }
    @keyframes float {
      0%,100% { transform: translateY(0); }
      50%     { transform: translateY(-10px); }
    }
  </style>
</head>
<body>
  <h1>Welcome to the Weird Zone</h1>

  <section class="card">
    <h2>Random Facts</h2>
    <p>
      The purple kangaroo flew over the neon desert, singing lullabies to
      the electric cactus. Meanwhile, a spoonful of clouds drifted by, 
      whispering secrets to the pixelated moon.
    </p>
  </section>

  <div class="gallery">
    <img src="https://picsum.photos/seed/alpha/300/200" alt="Random 1">
    <img src="https://picsum.photos/seed/bravo/300/200" alt="Random 2">
    <img src="https://picsum.photos/seed/charlie/300/200" alt="Random 3">
  </div>

  <section class="card">
    <h2>Lorum Ipsum Dolor</h2>
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit.  
      Curabitur vulputate nunc ut justo scelerisque, a dictum 
      lorem aliquet. Sed vehicula velit at dolor fermentum 
      consequat. Phasellus in libero vitae sapien blandit tincidunt.
    </p>
  </section>

  <footer>Everything is upside‑down and inside‑out 🌪️</footer>
      <script>
            window.__complyage__ = window.__complyage__ || {
                  apiKey      : '302b594b194d4b7ccea133eda5c8b1b8f7e85db2009185bead04971b9d2b574b',
                  baseUrl     : 'http://localhost:8089',
                  debug       : true,
                  version     : 1.01
            };
      </script>
      <script src="http://localhost:8089/static/js/loader.dev.js?cb=<?= time() ?>"></script>
</body>
</html>