# Building Words Counter Tool
from [https://codingchallenges.fyi](https://codingchallenges.fyi)

# step 0

Iniziando dalla  challenge https://codingchallenges.fyi/challenges/challenge-wc del contatore di parole ho implementato lo step zero.

Creato il main vuoto, un test vuoto, un file "wc.go" che conterrá tutta la logica, inizializzato il modulo "wc" e la cartella "target" con il compilato che attualemnte non fa niente.

# step 1

in questo step il tool "ccwc" deve restituire il numero di bytes in un file, utilizzando il file test.txt ci si aspetta (expected) 342190 bytes. 

il package `os` mette a disposizione il metodo `ReadFile` che restituisce un array di byte e va da se che contarli é semplice come contare la `len` di un array.

```golang
content, err := os.ReadFile(inputFile)
if err != nil {
    return nil, err
}
```
la prima cosa che ho creato cercando un approccio di tipo [TDD](https://it.wikipedia.org/wiki/Test_driven_development) creando il test `TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes` mettendoci dentro quello che mi aspetto e cercando di farlo compilare. In particolare ho creato una struttura `wc` con il suo costruttore `NewWc` che riceve il path del file di cui deve contare i bytes ed espone il metodo `CountBytes()`. In prima battuta il metodo ritornerá 0 bytes mentre nell'assert ce ne aspettiamo 342190 se utilizziamo il file `test.txt`... il test quindi fallirá.
```
--- FAIL: TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes (0.00s)
    c:\GoCode\src\github.com\diegoavanzini\learnfromcodechallenges\challenge-wc\wc_test.go:17: 
        	Error Trace:	c:/GoCode/src/github.com/diegoavanzini/learnfromcodechallenges/challenge-wc/wc_test.go:17
        	Error:      	Not equal: 
        	            	expected: 342190
        	            	actual  : 0
        	Test:       	TestWc_WhenInputFile_ShouldReturnTheExpectedNumberOfBytes
FAIL
FAIL	ccwc	1.183s
FAIL
```
vedi commit

vado ad implementare il metodo `CountBytes()` in modo da far passare il test
```
$ make test
go test .
ok      ccwc    1.006s
```
abbiamo il nostro primo test verde! 

Aggiungo il test che verifica, in caso di file errato, che il metodo ritorni un errore.

Infine é necessario chiamare il nostro eseguibile da linea di comando e far arrivare al nostro metodo i parametri che gli servono per fare quello che deve.

Per questo ci viene in aiuto il package `flag`.

In Go la funzione `main` é il punto di accesso del nostro tool ed é lei che riceve i parametri in input al tool.

Al fine di testare la validazione dei parametri in ingresso ho creato la funzione privata `validateInput` che riceve in ingresso i parametri passati e utilizzando il pacchetto `flag` definisce i parametri accettati e ne esegue il parsing e la validazione.

Il primo test verifica che senza parametri in ingresso la validazione ritorni un errore:

```golang
func TestWc_validateInputWithoutArguments_ShouldReturnError(t *testing.T) {
	// ARRANGE
	args := []string{}

	// ACT
	_, _, err := validateInput(args)
	defer resetFlags()

	// ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, "count what?", err.Error())
}
```

In particolare con 

```golang
byteCount = flag.CommandLine.Bool("c", false, "count bytes in file usage: ccwc -c <file>")
```
dichiaro che uno dei parametri del mio tool sará un booleano che si chiamerá "c" ed di default sará `false`, mentre se presente sará `true`.

con 
```golang
flag.CommandLine.Parse(args)
```
vado a fare il parsing degli argomenti.

Gli argomenti in ingresso saranno quelli passati a line di comando `os.Args[1:]` togliendo il primo che é il nome del comando oppure quelli passati nel test.

Da notare la chiamata in defer della funzione `resetFlags()`, questa é necessaria in quando la dichiarazione di `CommandLine` é statica, questo fa si che quando lancio il secondo test che va a ridefinire `c` questo vada in errore ma noi non lo vogliamo. Per evitarlo lanciamo questo 
```golang
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //flags are now reset
}
```

altri test che ho ritenuto utili sono quando il flag é presente ma non c´é il nome del file (`TestWc_validateInputWithCountFlagButNoFilepath_ShouldReturnError`) oppure c´é il nome del file ma é sbagliato (`TestWc_WhenInputFilePathIsWrong_ShouldReturnAnError`).


# step 2

in questo step dobbiamo calcolare il numero di righe in un file, utilizzando il file test.txt ci si aspetta (expected) 7145 righe.

Iniziamo come al solito dal test con il file di txt in dotazione e ritorno in prima battuta 0 facendo fallire il test.

Il test si aspetta il metodo `CountRows` associato alla struc `wc`, questo ritornerá il numero di linee. Per farlo deve anch'esso andare a leggere il file in input e vien da se che é necessario in questo caso entrare nel merito del contenuto del file per individuare gli a capo `[]byte{'\n'}`.

Sia che si tratti di leggere il numero di bytes o il numero di righe é necessaria la lettura del file. Per leggere il file si usa il descrittore ritornato da `os.Open(w.filepath)` e utilizzando il metodo `Read` andiamo a leggere e mettere il contenuto in un array di byte.

Centralizzo il calcolo in un metodo che si occupa di leggere il file e instanaziare i contatori che ci interessano.

```golang
func (w *wc) readFile(inputFile string) error {
	r, err := os.Open(w.filepath)
	if err != nil {
		return err
	}
	buf := make([]byte, 32*1024)
	for {
		c, err := r.Read(buf)
		w.numberOfLine += bytes.Count(buf[:c], []byte{'\n'})
		w.numberOfBytes += c
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
	}
}
```
ora é necessario introdurre il flag a linea di comando che dice di calcolare il numero di linee quindi altro test rosso che andiamo a mettere a posto.
E se inserisco piú flag ritorno  il conteggio di linee e bytes.


# step 3

in questo step dobbiamo calcolare il numero di parole nel testo. Il procedimento é il solito: test rosso per il nuovo conteggio => implementazione => flag verde e idem per il nuovo flag `-w`

Dopo vari tentativi ho trovato che invece di contare le parole sarebbe piú furbo e semplice contare il numero di spazi tra una parola e l'altra con una semplice espressione regolare `\s+` (cerca tutti i match con i spazi che si ripetono da una a piú volte conscutivamente)

Il test `TestWc_WhenInputFileAndLFlagL_ShouldReturnTheExpectedNumberOfLines` diventa verde.

Ora non ci resta che testare e implementare l'input con il flag `-w` ma a questo punto la funzione validateInput dovrebbe ritornare 3 flag e si rende necessario dare un significato a questo output raggruppando questi flag in una struttura.

Creo la struttura `WordCountInput` che contiene i 3 flag in ingresso al nostro tool e il `filepath` questo richiede un pó di refactoring ma mi permette di non modificare tutti gli utilizzatori del `validateFlag` ogni volta che aggiungo un nuovo flag.

# [step 4](./commit/81845ca)

in questo step si aggiunge il flag `-m` che serve per contare il numero di caratteri.

nella richiesta viene fatto presente che la risposta dipende dal proprio "locale", per golang ho trovato questa pagina [https://go.dev/blog/matchlang](https://go.dev/blog/matchlang) che sembra interessante... inoltre ho visto che esiste il modulo `utf8` che espone la funzione `utf8.RuneCountInString(str)` che ci viene in aiuto. Proviamolo creando prima il test per vedere se i conti tornano.

Test aggiunto, implementazione eseguita e sembra tutto ok, aggiungo il nuovo flag.

# [step 5](./commit/9748ce4)

in questo step si vuole che senza nessuna opzione ma solo il nome del file in ingresso vengano visualizzati i primi tre contatori come se passassi `-c` `-l` e `-v`.

La modifica per gestire il caso senza parametri ha ovviamente rotto il test che controllava che ci fosse almeno un parametro e in caso contrario si aspettava un errore.

É stata richiesta anche la modifica dell'output del tool. I test in questo caso sono a mano. 

# [step 6](./commit/9748ce4)

in questo step viene richiesto di poter concatenare il tool a linea di comando con altri tool, ricevendo in input il risultato del comando che lo precede con lo Unix pipe.
In Unix il pipe `|` é un costrutto molto potente che permette di redirigere l'output di un comando all'input del comando successivo, questo permette, pur utilizzando dei comandi semplici in sequenza, di creare workflow complessi.

Per capire se il nostro comando é lanciato con il pipe é ecessario eseguire questa "magia"
```golang
func IsPipe() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	modeChar := fi.Mode() & os.ModeCharDevice
	return modeChar == 0, nil
}
``` 

`os.Stdin` permette di leggere lo standard input e con il metodo `Stat()` possiamo avere informazioni su di esso, in particolare il Mode é un valore integer che combina diversi flags che descrivono i permessi e alcuni attributi in questo caso dello standard input. Il Mode é definito come una maschera di bit e ogni bit corrisponde a un determinato attributo o permesso.
L'operatore bitwise compara i bit di `fi.Mode()` uno ad uno con la costante `os.ModeCharDevice` che é `000001000000000000000000000`.

|& | pipe input | no pipe input| 
| ---- | ---- | ---- |
| os.ModeCharDevice| 000001000000000000000000000 | 00000`1`000000000000000000000 |
| fi.Mode() | 010000000000000000110110110 |10000`1`000000000000110110110|
| risultato | 000000000000000000000000000 |00000`1`000000000000000000000 |

Il valore del primo in caso di stdin valorizzato (pipe di unix) é `10000000000000000110110110` in questo caso il risultato del bitwise é `00000000` perché non ci sono due bit a 1 nella stessa posizione mentre se lo stdin non é valorizzato il `fi.Mode()` vale `100001000000000000110110110` abbiamo il bit in posizione 7 valorizzato a 1 in entrambi. Quindi quando il bitwise ritorna 0 (`modeChar == 0`) siamo nella situazione di pipe.

Ok ora che siamo in grado di capire se l'input arriva dalla pipeline o se dobbiamo prenderlo da un file é necessario capire come distinguere i due casi e comportarsi diversamente a seconda che siamo in un caso o nell'altro. Quello che cambia é il modo di leggere l'input.
Sento forte odore di strategy pattern... vediamo come implementarlo.

Mi viene spontaneo creare un `inputReader`, un'interfaccia che espone il metodo `Read()` e che verrá implementata dalla struct `FileInputReader` che si occupa di leggere il file e da `PipeInputReader` che si occupa di leggere dallo stdin. Il metodo Read ritorna un array di byte con il contenuto letto e un eventuale errore. Il costruttore esegue la validazione dei flag in ingresso, determina che tipo di input deve leggere e crea l'input reader corretto.
L' `inputReader` verrá passato poi nella creazione del tool che chiamando semplicemente il metodo Read() dell'implementazione corretta otterrá il contenuto ed eseguirá il conteggio dei contaori e in base ai flag in ingresso restituirá i valori calcolati al main che si preoccuperá di stamparli a video nel modo corretto in base ai flag impostati in ingresso.

Ho quindi introdotto il package reader con l'interfaccia e le relative implementazioni e costruttori e i test che prima testavano il `validateFlag` adesso verificano che i flag restituiti dall'inputReader siano gli stessi.
 
