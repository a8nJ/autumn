let m1 = {Sun: 10, Mon: 11};
// example 1
for (let s1 in m1) {
   console.log(s1);
}
// example 2
for (let [s1, n1] of Object.entries(m1)) {
   console.log(s1, n1);
}
