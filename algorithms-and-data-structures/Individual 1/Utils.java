package aed.loops;



public class Utils {
  

public static int maxNumRepeated(Integer[] a, Integer elem)  {
 
var indices = [];
var idx = a.indexOf(elem);
while (idx != -1) {
  indices.push(idx);
  idx = a.indexOf(elem, idx + 1);
}
return indices.length;
  
  }


}
