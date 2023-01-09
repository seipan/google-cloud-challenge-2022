import { Inter } from '@next/font/google';
import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { EventConfirmModal } from '../../../../components/molecules/Event/EventConfirmModal';

export default function Participate() {
  // ここは参加者がみるただの参加フォーム
  // TODO : SSRで実装してリンクを貼った時にOGPを表示させるようにする
  const [name, setName] = useState('');
  const [word, setWord] = useState('');
  const [isConfirm, setIsConfirm] = useState(false);
  const eventId = useRouter().query.eventId;
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        {isConfirm ? (
          <EventConfirmModal
            onParticipate={() => {}}
            onCancel={() => setIsConfirm(false)}
            eventId={Number(eventId as string)}
          />
        ) : (
          <>
            <EventInfo />
            <FormWrapper>
              <Input
                type="text"
                label="名前"
                content={name}
                changeContent={setName}
                required={true}
              />
              <Input
                type="text"
                label="ひとこと"
                content={word}
                changeContent={setWord}
                required={true}
              />
              <Button
                className="flex m-auto my-5"
                onClick={() => setIsConfirm(true)}
              >
                参加する
              </Button>
            </FormWrapper>
          </>
        )}
      </BasicTemplate>
    </>
  );
}
