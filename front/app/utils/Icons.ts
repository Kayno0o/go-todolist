import { IconDefinition } from '@fortawesome/fontawesome-svg-core';
import { faBook, faBriefcase, faHeadset, faHome, faPlaneDeparture } from '@fortawesome/free-solid-svg-icons';
import { faClipboard, faMap } from '@fortawesome/free-regular-svg-icons';

export type IconType = 'Clipboard' | 'Plane' | 'Headset' | 'Home' | 'Map' | 'Briefcase' | 'Book';

export const TaskIcons: { [name in IconType]: IconDefinition } = {
  Clipboard: faClipboard,
  Plane: faPlaneDeparture,
  Headset: faHeadset,
  Home: faHome,
  Map: faMap,
  Briefcase: faBriefcase,
  Book: faBook,
};
