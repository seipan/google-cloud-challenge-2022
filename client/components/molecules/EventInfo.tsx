import { Map } from '../atoms/map/Map';
import { EventWrapper } from './Event/EventWrapper';
import { Hanging } from './Event/Hanging';
import { EventBasicInfo } from './Event/EventBasicInfo';
import { UserIcon } from './User/UserIcon';
type TProps = {
  eventName: string;
  detail: string;
  location: string;
};
// 使う時になったら引数を入れる
export const EventInfo = () => {
  return (
    <>
      <Hanging />
      <EventWrapper>
        <UserIcon userName="miso" />
        <EventBasicInfo
          eventName="ラーメン行こう"
          detail="同志社周りのラーメン行こう！！あくたがわとかが良さげ"
        />
        <div className="lg:m-10 m-3">
          <Map lat={35.6809591} lng={139.7673068} />
        </div>
      </EventWrapper>
    </>
  );
};
