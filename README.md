# Lua Script API für App-Integration

Dieses Projekt erlaubt es, Lua-Skripte zur Erweiterung der App-Funktionalität zu schreiben.
Lua-Skripte werden automatisch aus dem Ordner `./scripts` geladen und können dynamisch
Einträge erzeugen und auf Benutzereingaben reagieren.

## Aufbau

### Verzeichnisstruktur

```
.
├── main.go
└── scripts/
    ├── beispiel.lua
    └── ...
```

### Ziele

- Lua-Skripte können dynamisch Items erstellen.
- Jedes Skript erhält Zugriff auf die globale Benutzereingabe.
- Aktionen können basierend auf dem ausgewählten Index durchgeführt werden.

---

## Lua API

### `getInput() -> string`

Gibt den aktuellen Benutzereingabetext zurück, der vom Nutzer eingegeben wurde.

```lua
local input = getInput()
```

---

### `addItem(title: string, subtitle: string) -> boolean`

Fügt ein Item mit Titel und Untertitel zur Ergebnisliste hinzu.

- Gibt `true` zurück, wenn dieses Item aktuell ausgewählt wurde (`ClickItemContainer()`).
- Andernfalls `false`.

```lua
local selected = addItem("Google", "Öffnet die Startseite")
if selected then
  -- Der Nutzer hat dieses Item ausgewählt
end
```

---

## Beispielskript (`scripts/example.lua`)

```lua
local input = getInput()

if input == "google" then
  local selected = addItem("Google öffnen", "https://www.google.com")
  if selected then
    os.execute("open https://www.google.com") -- macOS / passe für dein OS an
  end
elseif input == "github" then
  addItem("GitHub öffnen", "https://github.com")
end
```

---

## Integration in Go

Die App ruft bei Benutzereingabe `ParseInput(input)` auf und zeigt alle `addItem()`-Einträge an.
Beim Klicken auf ein Item wird `ClickItemContainer(index)` aufgerufen, wodurch dieselben
Lua-Skripte erneut mit gesetztem `selectedIndex` ausgeführt werden.

---

## Hinweise

- Nur `.lua`-Dateien im Ordner `./scripts` werden geladen.
- Alle Skripte werden bei jeder Eingabe und bei jedem Klick neu geladen.
- Die Reihenfolge der Items ist durch den Aufruf von `addItem` bestimmt.

---

## Anforderungen

- [Go](https://golang.org/)
- [`github.com/yuin/gopher-lua`](https://github.com/yuin/gopher-lua)

---

## Lizenz

MIT License
