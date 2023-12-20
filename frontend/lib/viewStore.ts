import { makeAutoObservable } from 'mobx';
import { View } from '@/types/index';

class ViewStore {
  view: View | null = null;

  constructor() {
    makeAutoObservable(this);
  }

  setView(view: View) {
    this.view = view;
  }
}

export default new ViewStore();
