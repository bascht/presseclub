# The (highly) unofficial Yakshed Presseclub

This is a small website wrapper around [SingleFile](https://github.com/gildas-lormeau/SingleFile) inspired by [@screenbreak/SingleFile-dockerized](https://github.com/screenbreak/SingleFile-dockerized) which exposes a HTTP endpoint that acts as a "transparent proxy".

The idea behind presseclub is to share paywalled articles with your friends.

## What presseclub does

- Presseclub is launched with a cookie store of the sites you or your friends have access to
- Every route behind `/lies/*` is downloaded through SingleFile with a headless Chromium and the supplied cookies
- Presseclub saves the file locally and makes the cache available at the very same URL

### Legal disclaimer

We have no idea if this is legal. (It's probably not).
It's fun though and you should set it up for your friends.

### Technical disclaimer

*DO NOT* export your complete cookie store for presseclub. The best way to retrieve the cookies you need for a specific paywalled site is to:

- Start Chromium with an empty / new Profile
- Browse to the site you want to add
- Use the [Get cookies.txt](https://chrome.google.com/webstore/detail/get-cookiestxt/bgaddhkoddajcdgocldbbfleckgcbcid?hl=en) extension to retrieve the Cookie Store


⚠️ Presseclub may be able make requests upon your behalf, so make sure to only share it with people you trust.
