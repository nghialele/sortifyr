import { API } from "./api";

export interface Track {
  id: number;
  spotifyId: string;
  name: string;
}

export const convertTrack = (t: API.Track): Track => {
  return {
    id: t.id,
    spotifyId: t.spotify_id,
    name: t.name,
  }
}

export interface TrackHistory extends Track {
  historyId: number;
  playedAt: Date;
}

export const convertTrackHistory = (h: API.TrackHistory): TrackHistory => {
  return {
    ...convertTrack(h),
    historyId: h.history_id,
    playedAt: new Date(h.played_at),
  }
}

export const convertTrackHistories = (h: API.TrackHistory[]): TrackHistory[] => {
  return h.map(convertTrackHistory)
}

