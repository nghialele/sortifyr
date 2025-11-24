import { API } from "./api";
import { convertUser, User } from "./user";

export interface Playlist {
  id: number;
  spotifyId: string;
  owner?: User;
  name: string;
  description?: string;
  public: boolean;
  tracks: number;
  collaborative: boolean;
  updatedAt: Date;
}

export const convertPlaylist = (playlist: API.Playlist): Playlist => {
  return {
    id: playlist.id,
    spotifyId: playlist.spotify_id,
    owner: playlist.owner ? convertUser(playlist.owner) : undefined,
    name: playlist.name,
    description: playlist.description,
    public: playlist.public,
    tracks: playlist.tracks,
    collaborative: playlist.collaborative,
    updatedAt: new Date(playlist.updated_at),
  }
}

export const convertPlaylists = (playlists: API.Playlist[]): Playlist[] => {
  return playlists.map(convertPlaylist)
}
