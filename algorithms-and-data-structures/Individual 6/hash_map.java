import java.util.HashMap;

public static void main(String [] args){
  HashMap <Integer, String> map = new HashMap<Integer, String>();

  for (Map.Entry<Integer, Sting> entry : map.entrySet()){
    if (entry.getKey() == 0){
      return entry.getValue();
    }
  }
  return map;
}

