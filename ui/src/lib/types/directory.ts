import { z } from "zod";
import { API } from "./api";
import { JSONBody } from "./general";
import { convertPlaylists, convertPlaylistsSchema, Playlist, playlistSchema } from "./playlist";
import { getUuid } from "../utils";

export interface Directory {
  id: number;
  name: string;
  children?: Directory[];
  playlists: Playlist[];
}

export const convertDirectory = (d: API.Directory): Directory => {
  return {
    id: d.id,
    name: d.name,
    children: d.children ? convertDirectories(d.children) : undefined,
    playlists: convertPlaylists(d.playlists),
  }
}

export const convertDirectories = (d: API.Directory[]): Directory[] => {
  return d.map(convertDirectory).sort((a, b) => a.name > b.name ? 1 : -1)
}

export const convertDirectorySchema = (directories: Directory[]): DirectorySchema[] => {
  return directories.map(d => ({
    id: d.id,
    iid: getUuid(),
    name: d.name,
    children: d.children ? convertDirectorySchema(d.children) : undefined,
    playlists: convertPlaylistsSchema(d.playlists),
  })).sort((a, b) => a.name > b.name ? 1 : -1)
}

export const directorySchema = z.object({
  id: z.number().optional(),
  iid: z.string().nonempty(),
  name: z.string(),
  get children() {
    return z.array(directorySchema).optional()
  },
  playlists: z.array(playlistSchema),
})
export type DirectorySchema = z.infer<typeof directorySchema> & JSONBody;
