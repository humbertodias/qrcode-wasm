// sw.js
self.addEventListener('install', function(event) {
    event.waitUntil(
        caches.open('qrcode-cache').then(function(cache) {
            return cache.addAll([
                '/',
                '/index.html',
                '/styles.css',
                '/script.js'
                // Add paths to other static assets as needed
            ]);
        })
    );
});

self.addEventListener('fetch', function(event) {
    event.respondWith(
        caches.match(event.request).then(function(response) {
            return response || fetch(event.request);
        })
    );
});
