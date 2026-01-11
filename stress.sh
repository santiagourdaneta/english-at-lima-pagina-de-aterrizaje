#!/bin/bash
URL="https://english-at-lima-pagina-de-aterrizaj.vercel.app/"
for i in {1..100}
do
   curl -s -o /dev/null -w "Solicitud #$i - Código: %{http_code}\n" $URL &
done
wait
echo "Prueba de carga finalizada."