package aed.individual4;

import java.util.Iterator;

import es.upm.aedlib.Pair;
import es.upm.aedlib.Position;
import es.upm.aedlib.positionlist.PositionList;

public class MultiSetListIterator<E> implements Iterator<E> {
  PositionList<Pair<E,Integer>> list;

  Position<Pair<E,Integer>> cursor;
  int counter;
  Position<Pair<E,Integer>> prevCursor;

  public MultiSetListIterator(PositionList<Pair<E,Integer>> list) {
    this.list = list;
    // ... anade cosas aqui
  }

  public boolean hasNext() {
    // hay que modificar
    return false;
  }

  public E next() {
    // hay que modificar
    return null;
  }
  
  public void remove() {
    // opcionalmente se puede modificar para obtener mas puntos
    // NO ES OBLIGATORIO
  }
}
