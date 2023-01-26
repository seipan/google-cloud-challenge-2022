import { useState } from 'react';
import { createNewEvent } from '../../../core/api/event/create';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../components/atoms/text/TypoWrapper';
import { Input } from '../../../components/atoms/form/Input';
import { Textarea } from '../../../components/atoms/form/Textarea';
import { FormWrapper } from '../../../components/atoms/form/FormWrapper';
import { Button } from '../../../components/atoms/text/Button';
import { MapForm } from '../../../components/atoms/form/MapForm';
import { useUserInfoStore } from '../../../store/userStore';
import { useRouter } from 'next/router';
import { TMapPosition } from '../../../components/atoms/map/MapBasicInfo';
import { useEffect } from 'react';
import { useNoticeStore } from '../../../store/noticeStore';

export default function New() {
  const router = useRouter();
  const [origin, setOrigin] = useState('');
  const { userInfo } = useUserInfoStore();
  const { changeNotice } = useNoticeStore();
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState<number>(1);
  const [detail, setDetail] = useState('');
  const [location, setLocation] = useState<null | TMapPosition>(null);
  useEffect(() => {
    setOrigin(window.location.origin);
  }, []);
  const createEvent = createNewEvent(
    (ok) => {
      router.push(`${origin}/event/${ok.created_event.event_id}/completion`);
      changeNotice({ type: 'Success', text: '作成に成功しました' });
    },
    (e) => {
      changeNotice({ type: 'Error', text: '作成に失敗しました' });
    }
  );
  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">募集する</h1>
      </TypoWrapper>

      <FormWrapper
        onSubmit={() => {
          createEvent(
            {
              user_id: userInfo.userId,
              user_name: 'atode', //TODO :
              icon: 'mada',
            },
            {
              event_name: name,
              max_member: Number(capacity),
              detail: detail,
              location: JSON.stringify(location),
              timestamp: Date.now(),
            }
          );
        }}
      >
        <Input
          type="text"
          label="イベント名"
          content={name}
          changeContent={setName}
          required={true}
          maxLength={50}
          minLength={1}
        />
        <Input
          type="number"
          label="募集人数(最大5名)"
          min={1}
          max={5}
          content={capacity}
          changeContent={setCapacity}
          required={true}
        />
        <Textarea
          label="詳細"
          content={detail}
          changeContent={setDetail}
          minLength={0}
          maxLength={300}
          required={true}
        />
        <MapForm location={location} setLocation={setLocation} />
        <Button className="flex m-auto my-5">募集する</Button>
      </FormWrapper>
    </BasicTemplate>
  );
}
