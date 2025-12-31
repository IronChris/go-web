# HTMX Cheatsheet

## Installation
```html
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
```

## Core Attribute

**hx-get/post/put/patch/delete** - Sendet HTTP-Request zur angegebenen URL
```html
<button hx-get="/api/data">Laden</button>
<button hx-post="/api/save">Speichern</button>
```

## Targeting & Swapping

**hx-target** - Bestimmt, welches Element aktualisiert wird
- `#id` - Element mit dieser ID
- `.class` - Erstes Element mit dieser Klasse
- `closest .class` - Nächster Eltern-Container
- `this` - Das Element selbst

```html
<button hx-get="/data" hx-target="#result">Laden</button>
<div id="result"></div>
```

**hx-swap** - Wie die Antwort eingefügt wird
- `innerHTML` - Ersetzt Inhalt (default)
- `outerHTML` - Ersetzt komplettes Element
- `beforebegin` - Vor dem Element
- `afterbegin` - Als erstes Kind
- `beforeend` - Als letztes Kind
- `afterend` - Nach dem Element
- `delete` - Löscht Zielelement
- `none` - Aktualisiert nichts

```html
<div hx-get="/more" hx-swap="beforeend">
  Bestehender Inhalt
</div>
```

## Trigger

**hx-trigger** - Wann Request ausgelöst wird
- `click` - Bei Klick (default für Buttons)
- `change` - Bei Änderung (default für Inputs)
- `submit` - Bei Submit (default für Forms)
- `load` - Beim Laden
- `revealed` - Wenn ins Viewport gescrollt
- `every 2s` - Alle 2 Sekunden

**Modifikatoren:**
- `once` - Nur einmal
- `changed` - Nur wenn Wert geändert
- `delay:1s` - Verzögerung
- `throttle:1s` - Maximal alle 1s

```html
<input hx-get="/search" hx-trigger="keyup changed delay:500ms">
<div hx-get="/news" hx-trigger="every 30s"></div>
<div hx-get="/data" hx-trigger="load, click"></div>
```

## Request Konfiguration

**hx-params** - Welche Parameter gesendet werden
- `*` - Alle (default)
- `none` - Keine
- `param1,param2` - Spezifische Parameter

**hx-vals** - Zusätzliche Werte als JSON
```html
<button hx-post="/api" hx-vals='{"userId": 123}'>Send</button>
```

**hx-headers** - Custom HTTP Headers
```html
<div hx-get="/data" hx-headers='{"Authorization": "Bearer token"}'>
```

**hx-include** - Zusätzliche Inputs einbeziehen
```html
<button hx-post="/submit" hx-include="[name='email']">Send</button>
```

## Indikatoren & UI

**hx-indicator** - Zeigt Loading-Indikator
```html
<button hx-get="/data" hx-indicator="#spinner">
  Laden
</button>
<img id="spinner" class="htmx-indicator" src="spinner.gif"/>
```

**CSS-Klassen:**
- `.htmx-request` - Während Request aktiv
- `.htmx-swapping` - Während Swap
- `.htmx-settling` - Während Settle-Phase

## Bestätigungen & Prompts

**hx-confirm** - Bestätigung vor Request
```html
<button hx-delete="/user/1" hx-confirm="Wirklich löschen?">
  Löschen
</button>
```

**hx-prompt** - Eingabe-Prompt
```html
<button hx-delete="/user" hx-prompt="Namen eingeben:">
  Löschen
</button>
```

## Erweiterte Features

**hx-push-url** - URL im Browser aktualisieren
```html
<a hx-get="/page2" hx-push-url="true">Seite 2</a>
```

**hx-select** - Nur Teil der Antwort verwenden
```html
<button hx-get="/page" hx-select="#content">
  Laden
</button>
```

**hx-boost** - Normale Links/Forms mit AJAX aufwerten
```html
<div hx-boost="true">
  <a href="/page2">Seite 2</a>
</div>
```

**hx-preserve** - Element bei Swap beibehalten
```html
<div hx-preserve="true">Bleibt erhalten</div>
```

## Response Headers (Server-seitig)

Der Server kann diese Headers setzen um Client-Verhalten zu steuern:

- `HX-Trigger` - Event triggern
- `HX-Redirect` - Browser-Redirect
- `HX-Refresh` - Seite neu laden
- `HX-Replace-Url` - URL ersetzen
- `HX-Push-Url` - URL pushen
- `HX-Retarget` - Anderes Ziel wählen
- `HX-Reswap` - Andere Swap-Methode

## Events & JavaScript

**HTMX Events hören:**
```javascript
document.body.addEventListener('htmx:afterSwap', (event) => {
  console.log('Content wurde geladen');
});
```

**Wichtige Events:**
- `htmx:beforeRequest`
- `htmx:afterRequest`
- `htmx:beforeSwap`
- `htmx:afterSwap`
- `htmx:load`

## Beispiel: Infinite Scroll
```html
<div>
  <div hx-get="/items?page=2" 
       hx-trigger="revealed" 
       hx-swap="afterend">
    <p>Mehr laden...</p>
  </div>
</div>
```

## Beispiel: Live Search
```html
<input type="search" 
       name="q"
       hx-get="/search" 
       hx-trigger="keyup changed delay:500ms"
       hx-target="#results">
<div id="results"></div>
```

## Beispiel: Delete mit Bestätigung
```html
<button hx-delete="/items/123"
        hx-confirm="Wirklich löschen?"
        hx-target="closest .item"
        hx-swap="outerHTML">
  🗑️ Löschen
</button>
```
