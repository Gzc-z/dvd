public class Main {
    public static void main(String[] args) {
        Song song = new Song(
                1,
                "Minha Música",
                "assets/Song1.wav"
        );

        Player player = new Player();


        player.loadSong(song);
        player.play();
        System.out.println("Estado: " + player.status());


        long duracaoMs = player.getClipDurationMs();
        if (duracaoMs > 0) {
            try {
                Thread.sleep(duracaoMs);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }



        player.stop();
    }
}