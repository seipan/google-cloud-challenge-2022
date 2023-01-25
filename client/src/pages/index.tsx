import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/text/TypoWrapper';
import style from '../styles/title.module.css';
import { LinkTo } from '../components/atoms/text/LinkTo';
import { Button } from '../components/atoms/text/Button';
export default function Home() {
  return (
    <>
      <MyHead title="すきーま" description="すきーまの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className={style.title}>すきーま</h1>
        </TypoWrapper>
        <div className="my-10 grid grid-cols-1 gap-2">
          <LinkTo href="/event/new" className="m-1">
            イベント募集ページへ
          </LinkTo>
          <LinkTo href="/event" className="m-1">
            イベント一覧ページへ
          </LinkTo>
          <Button onClick={() => console.log('ログアウト処理')} className="m-1">
            ログアウト
          </Button>
        </div>
      </BasicTemplate>
    </>
  );
}
