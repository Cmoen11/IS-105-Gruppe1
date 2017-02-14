# Oppgave 1

Formål: bli kjent med systemet sitt  :+1:

##### Hvor mange prosesser som kjører på din datamaskin?
Vi kan se hvor mange prosesser som kjører på Windows ved å bruke taskmgr, enten ved å kjøre programmet ved å skrive "taskmgr" i terminal, eller skrive "tasklist" for å få liste over alle prosesser i terminal.

På min maskin kjører 100 prosesser for øyeblikket.


##### Hvor mange prosesser som kjører på din virtuelle server i nettskyen?
For ubuntu skriver jeg "top" i terminal for å få opp tilsvarende oversikt som viser 147 prosesser.



##### Kan man gi et nøyaktig antall? Begrunn.
Da jeg tolker prosesser som "Tasks" (set av program instruksjoner som er lastet i minnet) så kjører min server 147 prosesser, som vises i oversikten.


##### Hvor mange av prosessene som “kjører”?
Av prosessene på min server er det bare 3 som kjører.


##### Hvis de ikke kjører, hvilke tilstander befinner de seg da?
Prosessene som ikke kjører er i tilstanden idle/sleep.


##### Hva er maskinvarespesifikasjon til din datamaskin (noter prosessortype, prosessorarkitektur, klokkefrekvens, informasjon om primært minne, størrelse på cache (både L1, L2 og L3 er ønskelig))?
Har laget en "ComputerInfo.go" fil for å hente ut denne informasjonen for windows / linux. 

- Lokal maskinvarespesifikasjoner:Prosessortype: Intel(R) Core(TM) i7-4770K , 4-core 8-threads
- Prosessorarkitektur: Intel`s Haswell, 64bit
- Klokkefrekvens:  4.1GHz 
- Informasjon om primært minne: 16GB DIMM DDR3 @ 1600 MHz
- Størrelse på cache ( 256 , 1024 , 8192)


##### Hvor mange CPU-“cores” har du tilgjengelig på din maskin? Noter.
Jeg har 4-Cores på min maskin som kjører HyperThreading teknologi


##### Hvor mange CPU-”cores” har du tilgjengelig på din virtuelle server? Noter.
Ved å bruke kommandoen "lscpu" ser jeg at jeg har 1 CPU med 1 Core på min virtuelle server.


##### Finn ut hvilken prosess i ditt system bruker mest minne. Beskrive denne prosessen kort.
Min datamaskin bruker mest minne på IntelliJ da jeg programmerer og kjører virtuelle prosesser med programmet.


##### Teamarbeid: Oppsummer alle data i en tabell i deres team-besvarelsen.
- Sammenlign deres platformer og diskuter forkjeller

<table style="height: 154px;" width="710">
<tbody>
<tr>
<td><strong>Navn</strong></td>
<td><strong>OS</strong></td>
<td><strong>CPU</strong></td>
<td><strong>RAM</strong></td>
<td><strong>HK</strong></td>
<td><strong>GPU</strong></td>
</tr>
<tr>
<td>Erlend W</td>
<td>Windows 10 pro</td>
<td>Intel i7 4770K, 3.5GHz</td>
<td>16GB 1333 MHz</td>
<td>Asus Z87 PRO Socket 1150</td>
<td>Nvidia 1070 MSI</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
<tr>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>
</tbody>
</table>

##### Hvilke komponenter (både fysiske og abstrakte) i deres datasystemer er involvert i oppstart, administrasjon og avslutning av prosesser? Definer komponentene du nevner
1. Power Supply : Når du trykker på POWER knappen, gir PSU strøm til hovedkortet og eksterne komponenter.
2. HK inneholder BIOS som sitter på grunnleggende instillinger for hvilke drivere som skal aktiveres og kjører en Power-On-Self-Test for så å starter bootloaderen.
3. BootLoaderen velger hvilken ressurs PCn skal starte fra (CDrom/HDD/USB)
4. Når bootloaderen finner OS, lastes dette inn på minnet og operativsystemet tar over oppstartsprosessen ved å laste inn sine drivere og instillinger.
5. Du kan nå logge inn på din PC.
