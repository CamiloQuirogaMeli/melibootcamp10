package main

/*
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus
servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
- un arreglo de números enteros con 100 valores
- un arreglo de números enteros con 1.000 valores
- un arreglo de números enteros con 1.000.000 valores

Para instanciar las variables utilizar rand

Se debe realizar el ordenamiento de cada una por:
- Ordenamiento por inserción
- Ordenamiento por burbuja
- Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1.000 y
después el de 1.000.000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber
qué ordenamiento fue mejor para cada arreglo
*/
