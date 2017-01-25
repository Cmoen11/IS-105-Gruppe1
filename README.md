# is105-uke04 : How to git?
<br> Step by step:
<br> Gå inn på mappen dere vil "lagre" Repoen på. (f.eks mkdir GitGroup  -> cd GitGroup)
<br> 1. Last ned med git clone kommandoen
<br> <i> git clone https://github.com/IS-105-GitGroup/is105-uke04.git</i>
<br>
<br> 2. Lag en ny branch med deres navn.
<br> <i> git checkout -b Erlend </i>
<br> Du ser nå at du er i branchen ved at branchnavn står etter mappestrukturen: GitGroup/is105-uke04 (Erlend)
<br>
<br> 3. Opprett en ny .go Fil i Go mappa enten med kommandolinje eller OS brukergrensesnitt.
<br> Når du har lagret filen kan du skrive: <i> git status </i> og vil se at filen er rød der.
<br>
<br> 4. Legg til filen, commit og last opp.
<br> <i> git add . </i> : "." gjør at du legger til alle filer, du kan også skrive filnavn.
<br> <i> git commit -m "kommentar"</i> : Commit filen med kommentar
<br> <i> git push origin Erlend </i> : Laster opp branchen Erlend til Repoen vår.
<br>
<br><b>git clone "url til repository"     </b>// clone et repository fra github f.eks.
<br><b>git status</b>				  // vil gi en oversikt over endrede filer
<br><b>git add "filename" 		  </b>// vil legge til filen til stage for commit
<br><b>git add .		  </b>// vil legge til alle filer som er endret til stage for commit
<br><b>git reset "filename"		</b>   // fjerne staged file for commit.
<br><b>git commit	-m “melding”	</b>  // pakke inn endringene klar til push
<br><b>git push origin "branchnavn" </b>// Legg ut commit’en til github.
<br><b>git checkout -b "branchnavn" </b>// opprette en ny branch
<br><b>git checkout "branchnavn" 	  </b>// vil gå til branch
<br><b>git merge "branchnavn"	</b>// merger branchen man er i til spesifisert branch. 
