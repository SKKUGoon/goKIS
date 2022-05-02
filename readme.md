# KIS Developers API with Golang

### Project 목표
<p> 
    Golang으로 쓰여진 KIS API를 구현.
    구현된 KIS API를 배포
</p>

## 주문 방식
### Hash Key
<p>
    Hash Key는 Request Body 의 값을 중간에 탈취하여 변조하지 못하도록 하는데 사용됨. 
    주문과 같은 POST 방식 사용시, JSON Body 와 함께, JSON Body의 Encoded Version 인 Hash Key가
    동시에 입력되어야 한다. 
</p>

<p>
    Hash Key를 요청하면 접근 토큰이 발급됨. 발급된 토큰은 일정 기간 유효하며,
    유효기간 만료 전에 폐기할 수 있음.(defer)로 하기 
</p>

## 국내주식주문
### 주식주문(현금)
<p>
    
</p>

### 주식주문(신용)
<em>NotImplemented</em>
<p>
    사유: 전 빚투 하고싶지 않아요
</p>

### 주식주문(정정취소)
<p>
    
</p>

### 주식정정취소가능주문조회
<p>

</p>

### 주식일별주문체결조회
<p>

</p>

### 주식잔고조회
<p>

</p>

### 매수가능조회
<p>

</p>

## 국내주식시세
## 국내주식실시간