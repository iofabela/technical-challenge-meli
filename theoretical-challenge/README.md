# Technical Challenge Meli - Theroetical

### Procesos, hilos y corrutinas

<ul>
    <li>Un caso en el que usarías procesos para resolver un problema y por qué.</li>
    <details>
    <summary><b>Answer</b></summary>
    <p>
        &emsp;&emsp; Los procesos son entidades independientes y no pueden compartir información entre ellos en ejecución. Esto significa que cada proceso tiene su propio espacio de memoria y puede ejecutar código diferente.
        <br>Entonces un caso podría ser en el entrenamiento de modelos donde cada proceso puede ejecutar un modelo diferente con configuraciones distintas de hiperparámetros. Cada entrenamiento es un proceso con su CPU y memoria independiente sin interferir con los demás, crucial para evitar conflictos o sobrecarga de memoria compartida. Si un proceso falla, los demás no se ven afectados.
    </p>
    </details>
    </li>
    <li>Un caso en el que usarías threads para resolver un problema y por qué.
    <details>
    <summary><b>Answer</b></summary>
    <p>
        &emsp;&emsp;Los threads son adecuadas para resolver problemas que requieren un alto uso en CPU y pueden beneficiarse del procesamiento en múltiples núcleos. Además, las tareas no dependen entre sí y puedan ejecutarse de manera independiente.
        <br>Un ejemplo de caso podría ser un problema que consiste en procesar grandes cantidades de datos, como aplicar un filtro a imágenes de alta resolución o realizar cálculos científicos complejos en un conjunto de datos masivo.
    </p>
    </details>
    </li>
    <li>Un caso en el que usarías corrutinas para resolver un problema y por qué.
    <details>
    <summary><b>Answer</b></summary>
    <p>
        &emsp;&emsp;Dado que las corrutinas son más adecuadas para tareas ligeras y en el manejo de múltiples tareas concurrentes de manera eficiente y estructurada, especialmente en problemas donde se bloquea el flujo haciendose un cuello de botella.
        <br>Un ejemplo de caso podría ser la descarga de datos desde multiples fuentes. Supongamos que se tiene que descargar datos desde varias API's o fuentes al mismo tiempo. Entonces, se puede utilizar las corrutinas para realizar las descargas (de manera concurrente) en espera de la respuesta de las API's y ejecutar otras tareas que se necesiten en ese momento.
    </p>
    </details>
    </li>
</ul>

### Optimización de recursos del sistema operativo

Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de ellos información en una API HTTP. ¿Cómo lo harías? Explicar.

<details>
    <summary><b>Answer</b></summary>
    <p>
       &emsp;&emsp;Primeramente, se debe tener ciertas consideraciones para evitar problemas de rendimiento. Como las limitaciones (restricciones) del API en la tasa de consultas (rate-limits), latencia de red y los recursos del sistema.
       <br><br>Una vez que se tiene en cuenta estos factores, se puede utilizar la técnica de <u>"divide y vencerás"</u> para optimizar el rendimiento.
       <b>Dividiendo los elementos en lotes más pequeños y procesandolos de forma controlada</b>. Esto no solo permite cumplir con los límites de solicitudes por segundo, sino que también reduce la probabilidad de sobrecargar el sistema y enfrentar bloqueos temporales.
       Entonces, <b>se manejarán las solicitudes con concurrencia/asincronía divido en múltiples hilos</b>. Si la API lo permite, se puede usar compresión en las respuestas para minimizar el tamaño de los datos transferidos.
       <br>Pero si el volumen sigue siendo muy grande para gestionarlo en un solo servidor, entonces se abren más instancias o procesos (dependiendo de la arquitectura del sistema). Distribuyendo las tareas entre múltiples máquinas.
       <br><br>Finalmente,es importante implementar mecanismos que aseguren la disponibilidad y rendimiento óptimo del sitema. Esto puede incluir la utilización de caché, balanceo de carga, monitoreo de rendimiento, y otros métodos para garantizar la eficiencia y robutez del sistema. De esta manera, se puede abordar el problema para garantizar que el sistema funcione de manera confiable y escalable.
    </p>
</details>

### Análisis de complejidad

- Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con
  complejidades $O(n^2)$, $O(n^3)$, $O(2n)$ y $O(n*log(n))$, respectivamente, ¿Cuál de los algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.<details>
    <summary><b>Answer</b></summary>
    <p>
    &emsp;&emsp;Considerando los algoritmos con comportamiento asintótico. Conforme crece <b><i>𝑛</i></b> en el tiempo de ejecución. Se <u>descarta el crecimiento exponencial</u> (algoritmo A) dado que se hace impráctico para la mayoría de los casos, aunque podría ser útil para entradas muy pequeñas, su desempeño rápidamente se vuelve ineficiente. Y favorecería el algoritmo de <u>n logaritmo de n (algoritmo D)</u>, dado que es más eficiente para grandes cantidades de datos. El crecimiento es más lento que cualquier otra opción, lo que garantiza un buen desempeño incluso cuando <b><i>𝑛</i></b> es considerablemente grande.
    </p>
  </details>

- Asume que dispones de dos bases de datos para utilizar en diferentes
problemas a resolver. La primera llamada <b>AlfaDB</b> tiene una complejidad de $O(1)$ en consulta y $O(n^2)$ en escritura. La segunda llamada <b>BetaDB</b> que tiene una complejidad de $O(log(n))$ tanto para consulta, como para escritura. ¿Describe en forma sucinta, qué casos de uso podrías atacar con cada una?<details>
  <summary><b>Answer</b></summary>
  <p>
  &emsp;&emsp;Para <b>AlfaDB</b> en casos dominados por consultas rápidas con datos estáticos y escrituras muy esporádicas, donde las consultas son frecuentes, pero los datos cambian rara vez. Como por ejemplo, tablas estáticas (ID de países, categorías, etc). 
  <br>Para <b>BetaDB</b> en casos dominados donde hay un volumen importante tanto de consultas como de escrituras. En entornos dinámicos donde se requiere un buen balance entre consulta y escritura. Como por ejemplo, bases de datos transaccionales y/o aplicaciones con actualizaciones frecuentes.
  </p>
</details>
