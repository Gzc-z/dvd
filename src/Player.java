import javax.sound.sampled.*;
import java.io.File;
import java.io.IOException;

public class Player {
    private String currentStatus;   //  "IDDLE" "PLAYING" "PAUSED"
    private boolean isPlaying;
    private Song currentSong;
    private Clip clip;

    public Player(){
        this.currentStatus = "IDLE";
        this.currentSong = null;
        this.isPlaying = false;
        this.clip = null;
    }

    public long getClipDurationMs() {
        if (clip != null && clip.isOpen()) {
            return clip.getMicrosecondLength() / 1000;
        }
        return 0;
    }


    public void loadSong(Song song){
        this.currentSong = song;
        this.isPlaying = false;
        this.currentStatus = "PAUSED";
        System.out.println("Música carregada: " + song.getName());
    }

    public void play(){
        if (currentSong == null) {
            currentStatus = "IDLE";
            System.out.println("Nenhuma música carregada!");
            return;
        }

        try {
            if (clip != null && clip.isOpen()) {
                clip.close();
            }

            File audioFile = new File(currentSong.getPath());

            AudioInputStream audioStream = AudioSystem.getAudioInputStream(audioFile);

            clip = AudioSystem.getClip();

            clip.open(audioStream);

            clip.start();

            isPlaying = true;
            currentStatus = "PLAYING";
            System.out.println("Tocando: " + currentSong.getName());

        } catch (UnsupportedAudioFileException e) {
            currentStatus = "IDLE";
            isPlaying = false;
            System.out.println("Formato não suportado. Use WAV.");


        } catch (IOException e) {
            currentStatus = "IDLE";
            isPlaying = false;
            System.out.println("Erro ao ler arquivo. Verifique o path: " + currentSong.getPath());

        } catch (LineUnavailableException e) {
            currentStatus = "IDLE";
            isPlaying = false;
            System.out.println("Sistema de áudio indisponível.");
        }
    }

    public void stop(){
        if (clip != null && clip.isRunning()) {
            clip.stop();
            clip.close();
        }

        isPlaying = false;

        if (currentSong == null) {
            currentStatus = "IDLE";
        } else {
            currentStatus = "PAUSED";
        }

        System.out.println("Música pausada.");
    }

    public boolean isPlaying() {
        return isPlaying;
    }

    public String status() {
        return currentStatus;
    }
}