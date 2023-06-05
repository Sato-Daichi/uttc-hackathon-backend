#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

$CMD_MYSQL -e "CREATE TABLE channels (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  create_user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  workspace_id varchar(255) NOT NULL
);"

$CMD_MYSQL -e "CREATE TABLE messages (
  id varchar(255) NOT NULL,
  text text NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE replies (
  id varchar(255) NOT NULL,
  text text NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  parent_id varchar(255) NOT NULL
);"

$CMD_MYSQL -e "CREATE TABLE users_channels (
  id varchar(255) NOT NULL,
  channel_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE users_workspaces (
  id varchar(255) NOT NULL,
  workspace_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE users (
  id varchar(26) NOT NULL,
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

$CMD_MYSQL -e "CREATE TABLE workspaces (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"

# ユーザーを作成
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000001', '佐藤', 'sato', 'sato@gmail.com');"
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000002', '鈴木', 'suzuki', 'suzuki@gmail.com');"
$CMD_MYSQL -e "INSERT INTO users (id, username, password, email) VALUES ('00000000000000000000000003', 'Tanaka', 'tanaka', 'tanaka@gmail.com');"

# ワークスペースを作成
$CMD_MYSQL -e "INSERT INTO workspaces (id, name) VALUES ('00000000000000000000000001', 'UTokyo Tech Club');"
$CMD_MYSQL -e "INSERT INTO workspaces (id, name) VALUES ('00000000000000000000000002', 'kaggler-ja');"

# チャンネルを作成
$CMD_MYSQL -e "INSERT INTO channels (id, name, description, create_user_id, workspace_id) VALUES ('00000000000000000000000001', 'general', '一般的な内容', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO channels (id, name, description, create_user_id, workspace_id) VALUES ('00000000000000000000000002', 'random', '雑談', '00000000000000000000000002', '00000000000000000000000001');"

# メッセージを作成
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000001', '山口県柳井市の市立小学校で、６年生の担任教諭が自身のタブレットに書き込んでいた学習状況などに関するメモを一部の児童が見たことで、「学校に行きたくない」などと訴えていたことが、校長や市教育委員会への取材で分かった。不適切な表現があったといい、学校は５日から当面、担任を別の教諭に代行させる措置を取った。校長らによると、２日の朝の会の前、打ち合わせのため担任が教師用タブレットを教壇に置いたまま約２０分間、教室を離れた。タブレットは操作ができる状態で、約３０人のクラスの半数ほどがメモを見た。メモは児童の学習への取り組みや学校内の人間関係、保護者の要望などを担任が個人的にまとめたもので、校長らが確認したところ、一部の児童について努力の成果が見られないとの趣旨を記した部分や、保護者とのやりとりに関する部分に不適切な表現があった。５日は１人が体調不良で欠席したという。市教委のタブレット利用規定では、教師用タブレットに児童の個人情報を入れることは禁止されていたという。校長は「児童や保護者に不快な思いをさせ、申し訳ない」と謝罪した。', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000002', 'ウクライナに侵略するロシアの首都モスクワで、妨害電波を出して無人機を無力化する銃や飛来を覚知する装置の需要が急増している。国境から約４５０キロ・メートル離れていながら５月末に無人機で攻撃され、安全地帯でなくなったとの危機感が芽生えたようだ。露有力紙コメルサントやタス通信によると、少なくとも８機の無人機がモスクワ方面に飛来した５月３０日以降、企業を中心に対無人機銃の引き合いが増えているという。', '00000000000000000000000001', '00000000000000000000000002');"
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000003', '５日午後４時２０分頃、東京都町田市本町田で、同市の無職男性（８２）の乗用車が歩道に乗り上げ、６０歳代とみられる歩行者の男性をはねて、そのまま東京消防庁町田消防署の外壁に突っ込んだ。はねられた男性は搬送先の病院で死亡した。無職男性は手に軽傷。警視庁町田署幹部によると、無職男性は現場近くのドラッグストアの駐車場から車を発進させた際に運転操作を誤ったとみられ、「ブレーキとアクセルを踏み間違えた」と説明している。同署が事故の状況を調べている。消防業務への影響はなかった。乗用車が突っ込んだ町田消防署の事故現場を調べる警察官ら（５日午後、東京都町田市で、読売ヘリから）＝上甲鉄撮影現場はＪＲ町田駅の北西約２・５キロで、住宅や店舗が立ち並ぶ地域。', '00000000000000000000000001', '00000000000000000000000003');"

$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000004', 'IT企業「Carelogy（ケアロジー）」（東京）は、猫がけがや病気で痛みを感じているか人工知能（AI）で検知するアプリを開発した。猫は表情やしぐさから痛みの有無を判断しにくく、飼い主が気づいたときには重症になっているケースもある。アプリが早期診断や治療に役立つと期待できるという。', '00000000000000000000000002', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000005', '動画で芸能人らを脅すなどしたとして、警視庁は４日、前参院議員のガーシーこと東谷義和容疑者（51）を暴力行為等処罰法違反（常習的脅迫）容疑などで逮捕した。東谷容疑者は滞在先のアラブ首長国連邦（ＵＡＥ）で事実上の国外退去処分を受けたとみられ、同日夕に航空機で帰国した。発表によると、東谷容疑者は昨年２～８月、芸能人と実業家、デザイナーの男性３人に対し、ユーチューブで配信した動画内で脅迫や名誉 毀損きそん にあたる発言を繰り返し、デザイナーに対しては事業からの撤退を強要するなどした疑い。', '00000000000000000000000002', '00000000000000000000000002');"
$CMD_MYSQL -e "INSERT INTO messages (id, text, channel_id, user_id) VALUES ('00000000000000000000000006', '出席者が乾杯のあいさつで「（いずれ来る）衆院選の勝利に向けて」と発する前に、首相は「ああ、いずれ来るね」と相づちを打ったという。出席者の一人は、「解散は常に総理の頭の中にはあるのだなと感じた」と振り返る。前回２０２１年衆院選では、全国に４７ある１区のうち、自民が「２９勝１８敗」で勝ち越した。政権復帰した１２年衆院選以降は、自民の堅調な戦いが目立つ。, '00000000000000000000000002', '00000000000000000000000003');"

# チャンネルとユーザーを紐付け
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000001', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000002', '00000000000000000000000001', '00000000000000000000000002');"
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000003', '00000000000000000000000001', '00000000000000000000000003');"

$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000004', '00000000000000000000000002', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000005', '00000000000000000000000002', '00000000000000000000000002');"
$CMD_MYSQL -e "INSERT INTO users_channels (id, channel_id, user_id) VALUES ('00000000000000000000000006', '00000000000000000000000002', '00000000000000000000000003');"

# ワークスペースとユーザーを紐付け
$CMD_MYSQL -e "INSERT INTO users_workspaces (id, workspace_id, user_id) VALUES ('00000000000000000000000001', '00000000000000000000000001', '00000000000000000000000001');"
$CMD_MYSQL -e "INSERT INTO users_workspaces (id, workspace_id, user_id) VALUES ('00000000000000000000000002', '00000000000000000000000001', '00000000000000000000000002');"
$CMD_MYSQL -e "INSERT INTO users_workspaces (id, workspace_id, user_id) VALUES ('00000000000000000000000003', '00000000000000000000000001', '00000000000000000000000003');"
