# Zadania

Wszystkie zadania polegają na napisaniu programu od nowa lub modyfikacji
instniejącego już kodu.

Słowem wstępu, wiele ćwiczeń jest prostych, często da się zadanie zrobić
w sposób bardziej zwięzły niż wymaga tego polecenie.
Zalecam trzymania się poleceń, w celu zapamiętania i przećwiczenia 
elementów składni języka.

## Zadania wstępne

### 1. Imię 
Napisz program, który wypisze "Hello <Twoje imię>!".  
Zainicjalizuj (stwórz) zmienną, do której najpierw przypiszesz swoje imię, a następnie
ją wypisz.

### 2. Kwadrat 
Napisz program, który policzy pole kwadratu o boku 30.03. 
Użyj w tym celu zmiennej lub stałej.   
Dla ciekawskich, można sformatować liczby przy ich wypisywaniu, czyli
nadać im określony format, np. wypisać tylko 2 miejsca po przecinku, 
w tym celu korzysta się z funkcji: `fmt.Printf()`. Funkcja ta jako pierwszy 
argument przyjmuje format jakiego chcemy użyć. Dla przykładu, kod:  
```go
var x int = 12.345
fmt.Println("%f", x)    // wypisze: 12.345
fmt.Println("%.2f", x)  // wypisze: 12.34
``` 

### 3. Pytanie o imię 
Przerób kod z ćwiczenia pierwszego, tak aby najpierw zapytał użytkownika
o jego imie, a następnie się przywitał (np. wypisał "Hello Paweł!").
Przykładowa interakcja z programem powinna wyglądać jak poniżej:
```
> Jak masz na imię?
> : Paweł
> Witaj Paweł!
``` 
    
### 4. Wiek  
Napisz program, który zapyta użytkownika o rok urodzenia 
(sprawdź `step4a.go`), a następnie policzy ile lat ma użytkownik.  
Do obliczeń na czasie w Go najłatwiej skorzystać z biblioteki
`time`. Aby pobrać do zmiennej obecny rok, można napisać:
```go
currentYear := time.Now().Year()
```  
Powyższa linijka, po rozbiciu na części składowe:  
* `time.Now()` zwraca obecny czas (znacznik czasu, zawiera także datę)
* `time.Now().Year()` zwraca obecny rok.

Masz już wszystko, czego trzeba aby wypisać wiek użytkownika.

A co jeżeli dzisiaj jest 16 stycznia, a użytkownik urodził się w grudniu?
Możesz dokonać dokładniejszych obliczeń ;-)

### 4. Pełnoletność
Wiek użytkownika jest informacją pochodną, tzn. ciągle się zmienia.
Skutkuje to tym, że co rok staje się nieaktualny. Dobry system 
zapyta użytkownika nie o wiek, a o rok urodzenia i na tej informacji
będzie opierał swoje działanie.

Dlatego, dokonaj modyfikacji programu z zadania 4. Wiek, tak, aby jeżeli 
użytkownik ma co najmniej 21 lat, to wyświetlił informację: 
"Dostęp przyznany", w przeciwnym przypadku, wyświetl użytkownikowi informację
"Brak dostępu".

Możesz też spróbować napisać program, który wypisze komunikat tylko dla
użytkownikow, których wiek jest liczbą nieparzystą.

### 5. Suma 5 liczb. 
Napisz program, który zapyta użytkownika o 5 liczb, a następnie wypisze ich
sumę.

Napisz dwie wersje programu. W pierwszej nie korzystaj z pętli, a w drugiej 
tak.

### 6. Suma zadanej liczby liczb.
Napisz program, który najpierw zapyta użytkownika o to ile liczb chce
wprowadzić, następnie będzie je wczytywał i na koniec wypisze sumę tych liczb.

### 7. Ciąg Fibonacciego
To klasyczne zadanie programistyczne polega na wypisaniu n-tej liczby 
Fibonacciego.

Początek tego ciągu wygląda następująco: 
1, 1, 2, 3, 5, 8, 13, 21, 24, 45,...

Pierwsze dwa elementy tego ciągu, są zdefiniowane jako:
* Pierwszy element ciągu to 1, oznaczamy go przez `fib(1)`.
* Drugi element ciągu to także 1, oznaczamy go przez `fib(2)`.
* Kolejny element ciągu jest równy sumie dwóch poprzednich. Tzn. że
trzeci element to suma pierwszego i drugiego.
Możemy to zapisać jako `fib(3)=fib(1)+fib(2)`.
W ogólności, możemy napisać, że: `fib(n)=fib(n-1)+fib(n-2)`

1. Twoim celem w tym ćwiczeniu jest napisać program, który wypisze pierwsze
15 liczb Fibonacciego. Wykorzystaj do tego pętlę.
2. Następnie dokonaj zmian w programie. Zapytaj użytkownika, którą liczbę
Fibonacciego chciałby zobaczyć, po czym wypisz liczbę, o którą zapytał
użytkownik.
3. Napisz funkcję, która zwraca n-tą liczbę z ciągu Fibonacciego.
Możesz skorzystać z poniższej deklaracji funkcjis:
```go
func fib(n int) int
```