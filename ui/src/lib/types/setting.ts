import { API } from "./api";

export interface Setting {
  id: number;
  lastUpdated: Date;
}

export const convertSetting = (s: API.Setting): Setting => {
  return {
    id: s.id,
    lastUpdated: new Date(s.last_updated),
  }
}
