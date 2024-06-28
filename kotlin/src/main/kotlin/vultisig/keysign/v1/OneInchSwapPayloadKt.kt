// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: vultisig/keysign/v1/1inch_swap_payload.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package vultisig.keysign.v1;

@kotlin.jvm.JvmName("-initializeoneInchSwapPayload")
public inline fun oneInchSwapPayload(block: vultisig.keysign.v1.OneInchSwapPayloadKt.Dsl.() -> kotlin.Unit): vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload =
  vultisig.keysign.v1.OneInchSwapPayloadKt.Dsl._create(vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `vultisig.keysign.v1.OneInchSwapPayload`
 */
public object OneInchSwapPayloadKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload = _builder.build()

    /**
     * `.vultisig.keysign.v1.Coin from_coin = 1 [json_name = "fromCoin"];`
     */
    public var fromCoin: vultisig.keysign.v1.CoinOuterClass.Coin
      @JvmName("getFromCoin")
      get() = _builder.getFromCoin()
      @JvmName("setFromCoin")
      set(value) {
        _builder.setFromCoin(value)
      }
    /**
     * `.vultisig.keysign.v1.Coin from_coin = 1 [json_name = "fromCoin"];`
     */
    public fun clearFromCoin() {
      _builder.clearFromCoin()
    }
    /**
     * `.vultisig.keysign.v1.Coin from_coin = 1 [json_name = "fromCoin"];`
     * @return Whether the fromCoin field is set.
     */
    public fun hasFromCoin(): kotlin.Boolean {
      return _builder.hasFromCoin()
    }
    public val OneInchSwapPayloadKt.Dsl.fromCoinOrNull: vultisig.keysign.v1.CoinOuterClass.Coin?
      get() = _builder.fromCoinOrNull

    /**
     * `.vultisig.keysign.v1.Coin to_coin = 2 [json_name = "toCoin"];`
     */
    public var toCoin: vultisig.keysign.v1.CoinOuterClass.Coin
      @JvmName("getToCoin")
      get() = _builder.getToCoin()
      @JvmName("setToCoin")
      set(value) {
        _builder.setToCoin(value)
      }
    /**
     * `.vultisig.keysign.v1.Coin to_coin = 2 [json_name = "toCoin"];`
     */
    public fun clearToCoin() {
      _builder.clearToCoin()
    }
    /**
     * `.vultisig.keysign.v1.Coin to_coin = 2 [json_name = "toCoin"];`
     * @return Whether the toCoin field is set.
     */
    public fun hasToCoin(): kotlin.Boolean {
      return _builder.hasToCoin()
    }
    public val OneInchSwapPayloadKt.Dsl.toCoinOrNull: vultisig.keysign.v1.CoinOuterClass.Coin?
      get() = _builder.toCoinOrNull

    /**
     * `string from_amount = 3 [json_name = "fromAmount"];`
     */
    public var fromAmount: kotlin.String
      @JvmName("getFromAmount")
      get() = _builder.getFromAmount()
      @JvmName("setFromAmount")
      set(value) {
        _builder.setFromAmount(value)
      }
    /**
     * `string from_amount = 3 [json_name = "fromAmount"];`
     */
    public fun clearFromAmount() {
      _builder.clearFromAmount()
    }

    /**
     * `string to_amount_decimal = 4 [json_name = "toAmountDecimal"];`
     */
    public var toAmountDecimal: kotlin.String
      @JvmName("getToAmountDecimal")
      get() = _builder.getToAmountDecimal()
      @JvmName("setToAmountDecimal")
      set(value) {
        _builder.setToAmountDecimal(value)
      }
    /**
     * `string to_amount_decimal = 4 [json_name = "toAmountDecimal"];`
     */
    public fun clearToAmountDecimal() {
      _builder.clearToAmountDecimal()
    }

    /**
     * `.vultisig.keysign.v1.OneInchQuote quote = 5 [json_name = "quote"];`
     */
    public var quote: vultisig.keysign.v1.1InchSwapPayload.OneInchQuote
      @JvmName("getQuote")
      get() = _builder.getQuote()
      @JvmName("setQuote")
      set(value) {
        _builder.setQuote(value)
      }
    /**
     * `.vultisig.keysign.v1.OneInchQuote quote = 5 [json_name = "quote"];`
     */
    public fun clearQuote() {
      _builder.clearQuote()
    }
    /**
     * `.vultisig.keysign.v1.OneInchQuote quote = 5 [json_name = "quote"];`
     * @return Whether the quote field is set.
     */
    public fun hasQuote(): kotlin.Boolean {
      return _builder.hasQuote()
    }
    public val OneInchSwapPayloadKt.Dsl.quoteOrNull: vultisig.keysign.v1.1InchSwapPayload.OneInchQuote?
      get() = _builder.quoteOrNull
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload.copy(block: `vultisig.keysign.v1`.OneInchSwapPayloadKt.Dsl.() -> kotlin.Unit): vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayload =
  `vultisig.keysign.v1`.OneInchSwapPayloadKt.Dsl._create(this.toBuilder()).apply { block() }._build()

public val vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayloadOrBuilder.fromCoinOrNull: vultisig.keysign.v1.CoinOuterClass.Coin?
  get() = if (hasFromCoin()) getFromCoin() else null

public val vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayloadOrBuilder.toCoinOrNull: vultisig.keysign.v1.CoinOuterClass.Coin?
  get() = if (hasToCoin()) getToCoin() else null

public val vultisig.keysign.v1.1InchSwapPayload.OneInchSwapPayloadOrBuilder.quoteOrNull: vultisig.keysign.v1.1InchSwapPayload.OneInchQuote?
  get() = if (hasQuote()) getQuote() else null
