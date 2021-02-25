package aed.multisets;

import es.upm.aedlib.Pair;
import es.upm.aedlib.Position;
import es.upm.aedlib.positionlist.PositionList;
import es.upm.aedlib.positionlist.NodePositionList;


/**
 * An implementation of a multiset using a positionlist.
 */
public class MultiSetList<Element> implements MultiSet<Element> {

    /**
     * Datastructure for storing the multiset.
     */
    private PositionList<Pair<Element,Integer>> elements;

    private int size;


    /**
     * Constructs an empty multiset.
     */
    public MultiSetList() {
	this.elements = new NodePositionList<Pair<Element,Integer>>();
    }
}
