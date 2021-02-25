package aed.hashtables;

import es.upm.aedlib.Position;
import es.upm.aedlib.Entry;
import es.upm.aedlib.positionlist.*;
import es.upm.aedlib.map.Map;


/**
 * Implements the Map interface using a hashtable with separate chaining.
 */
public class HashTable<K,V> implements Map<K,V> {
    private NodePositionList<Entry<K,V>>[] table;
    private int size;

    /**
     * Creates a hash table with the (not modifiable) 
     * number of buckets specified by the <code>tableSize</code> parameter.
     */
    @SuppressWarnings("unchecked")
    public HashTable(int tableSize) {
	NodePositionList<Entry<K,V>>[] table =
	    (NodePositionList<Entry<K,V>>[])
	    new NodePositionList<?>[tableSize];
	for (int i=0; i<tableSize; i++)
	    table[i] = new NodePositionList<Entry<K,V>>();
	this.table = table;
    }

    private int index(Object key) {
	return table.get(key);
        // CHANGE
	//  index(key) == abs(key.hashCode()) modulo tamano(table)
	return 0;
    }

    /**
     * Is the table empty?
     */
    public boolean isEmpty() {
        boolean estaVacia = true;
        int i=0; 
	while (i<tableSize && estaVacia){
        if(!table[i].isEmpty()){
        estaVacia=false;}
        i++;}
        // CHANGE
	return estaVacia;
    }

    /**
     * Returns the size (the number of entries) of the table.
     */
    public int size() {
        if(!table.isEmpty()){
	Pointer <E> cursor = table[0].first();
        int contador = 0;
        for(int i=0; i<tableSize;i++){
        while (cursor!=null){
        contador++;
        cursor = table[i].next(cursor);}}}
	return contador;
    }

    /**
     * Checks if the map contains an entry with key.
     * Returns true if the map contains an entry with key, and
     * false otherwise.
     */
  public boolean containsKey(Object key) {
	boolean containsKey = false;
        int i=0;
        while (!containsKey && i<tableSize){
        containsKey = table[i].containsKey();
        i++;}
	return containsKey;
  }

    /**
     * Stores a key-value entry.
     * Returns the value of the previous entry with the specified key,
     * or null if no such entry existed.
     *
     * inutil public boolean coincide (K key,V value){
        *int i=0;
        *boolean encontrado=false;
        *boolean res = false;
        *V aux;
        *while(i<tableSize && !encontrado){
        *aux = table[i].get(key);
        *if (aux!=null){
        *encontrado = true;
        *res = aux.equals(value);}
        *i++;}
       * return res;
  */}
        
    public V put(K key, V value) 
throws InvalidKeyException{
        if(!table.containsKey(key)){
        throw new InvalidKeyException ();}
        int i=0;
        V value = null;
        while(i<tableSize){
        if(table[i].containsKey()){
        table[i].put(key,value);}
        i++;
        return value;
    }

    /**
     * Returns the value associated with the specified key,
     * or null if there is no entry with the key in the map.
     */
    public V get(K key) {
	V value = null;
        if (!table.containsKey(key)){
        return value;}
        int i=0;
        while(i<tableSize){
        if(table[i].containsKey()){
        table[i].get(key);}
        i++;}
    }

    /**
     * Removes the key-value entry with the specified key.
     * Returns the value of the removed entry, or null if no such entry
     * existed.
     */
    public V remove(K key) {
	V value = null;
        if (!table.containsKey(key)){
        return value;}
        int i=0;
        while(i<tableSize){
        if(table[i].containsKey()){
        table[i].remove(key);}
        i++;};
    }

    /**
     * Returns an iterable object containing all the keys in entries in
     * the map.
     */
    public Iterable<K> keys() {
	Iterable<K>res= new Iterable<K>();
        if (!table.isEmpty()){
        for(int i=0;i<tableSize;i++){
        Pointer<E>cursor = table[i].first();
        while(cursor!=null){
        res.next()=cursor.getElement().getLeft();
        cursor = table[i].next(cursor);}}
	return res;
    }

    /**
     * Returns an iterable object containing all the entries in
     * the map.
     */
     public Iterable<Entry<K,V>> entries() {
	Iterable<Entry<K,V>>res= new Iterable<Entry<K,V>>();
        if (!table.isEmpty()){
        for(int i=0;i<tableSize;i++){
        Pointer<E>cursor = table[i].first();
        while(cursor!=null){
        res.next()=cursor.getElement();
        cursor = table[i].next(cursor);}}
	return res;
    }

}
