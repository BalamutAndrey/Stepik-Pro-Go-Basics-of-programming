## 3.1.2. Пароль

При регистрации на сайтах требуется вводить пароль дважды. Это сделано для безопасности, поскольку такой подход уменьшает возможность неверного ввода пароля.

Напишите программу, которая сравнивает пароль и его подтверждение. Если они совпадают, то программа выводит: «**Пароль принят**», иначе: «**Пароль не принят**».

**Формат входных данных**
* На вход программе подаются две строки.

**Формат выходных данных**
* Программа должна вывести одну строку в соответствии с условием задачи.

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input 1:**
> **qwerty<br />
> qwerty**

**Sample Output 1:**
> **Пароль принят**

<br />

**Sample Input 2:**
> **qwerty<br />
> Qwerty**

**Sample Output 2:**
> **Пароль не принят**

<br />

**Sample Input 3:**
> **PythonROCKS<br />
> PythonROCKS**

**Sample Output 3:**
> **Пароль принят**
___
```Go
package main

func main() {
    // put your code here
}
```