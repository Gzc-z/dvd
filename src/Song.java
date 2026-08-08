public class Song {
    private int id;
    private String name;
    private String path;

    public Song(int id, String name, String path){
        this.id = id;
        this.name = name;
        this.path = path;
    }

    public String getName(){
        return name;
    }
    public int getId(){
        return id;
    }
    public String getPath(){
        return path;
    }
}
