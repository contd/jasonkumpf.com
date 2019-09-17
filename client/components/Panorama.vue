<template>
  <div
    class="panorama pv-container pv-animating"
    :class="{'pv-vertival':direction === 'vertical'}"
    :style="styling"
    @touchstart="touchHandler"
    @touchmove="touchHandler"
    @touchend="touchHandler"
    @mousedown="handle"
    @mouseup="handle"
    @mouseleave="handle"
  >
    <div v-if="caption && ready" class="pv-caption" :style="captionStyling">
      {{ caption }}
    </div>
    <div v-if="!ready" class="pv-loader" :style="loaderStyling">
      Loading...
    </div>
  </div>
</template>

<script>
const props = {
  source: {
    type: String,
    default: null
  },
  caption: {
    type: String,
    default: null
  },
  repeat: {
    type: Boolean,
    default: true
  },
  direction: {
    type: String,
    default: 'horizontal'
  },
  animationTime: {
    type: Number,
    default: 700
  },
  easing: {
    type: String,
    default: 'ease-out'
  }
}

export default {
  name: 'Panorama',
  props,
  data () {
    return {
      ready: false,
      width: 200,
      height: 100,
      imageWidth: 200,
      imageHeight: 100,
      bounds: {}
    }
  },
  computed: {
    styling () {
      let repeat = 'no-repeat'
      if (this.repeat === true) {
        repeat = 'repeat'
      }
      return {
        backgroundImage: `url(${this.source}) `,
        backgroundRepeat: repeat,
        transition: `background ${this.animationTime}ms ${this.easing}`,
        position: 'relative',
        float: 'left',
        width: '100%',
        height: '100%',
        minHeight: '100px',
        backgroundSize: `auto ${this.height}px`,
        cursor: 'move'
      }
    },
    loaderStyling () {
      return {
        position: 'absolute',
        top: '50%',
        left: '50%',
        marginLeft: '-30px',
        marginRight: '-9px',
        backgroundColor: 'white'
      }
    },
    captionStyling () {
      return {
        padding: '10px 10px',
        width: '100%',
        background: 'rgba(0,0,0,0.2)',
        color: 'white',
        textAlign: 'center',
        position: 'absolute',
        bottom: 0
      }
    }
  },
  mounted () {
    this.init()
  },
  beforeDestroy () {},
  methods: {
    init () {
      if (!this.source) {
        console.error('Image source is not defined!')
        return
      }
      const scope = this

      this.width = this.$el.clientWidth
      this.height = this.$el.clientHeight

      this.image = new Image()
      this.image.src = this.source
      this.image.onload = () => {
        scope.imageWidth = scope.image.naturalWidth
        scope.imageHeight = scope.image.naturalHeight
        const el = scope.$el
        const elbounds = {
          w: parseInt(el.clientWidth),
          h: parseInt(el.clientHeight)
        }
        this.bounds = {
          w: this.width - elbounds.w,
          h: this.height - elbounds.h
        }
        this.origin = {
          x: 0,
          y: 0
        }
        this.start = {
          x: 0,
          y: 0
        }
        this.movecontinue = false

        scope.ready = true
        scope.$emit('load')
      }
    },
    touchHandler (event) {
      const touches = event.changedTouches
      const first = touches[0]
      let type = ''

      switch (event.type) {
        case 'touchstart':
          type = 'mousedown'
          break
        case 'touchmove':
          type = 'mousemove'
          break
        case 'touchend':
          type = 'mouseup'
          break
        default:
          return
      }
      const simulatedEvent = document.createEvent('MouseEvent')
      let mult = 2

      if (navigator.userAgent.match(/Android/i)) {
        mult = 10
      }

      simulatedEvent.initMouseEvent(
        type,
        true,
        true,
        window,
        1,
        first.screenX,
        first.screenY,
        first.clientX * mult,
        first.clientY * mult,
        false,
        false,
        false,
        false,
        0,
        null
      )
      first.target.dispatchEvent(simulatedEvent)
    },
    move (e) {
      const inbounds = {
        x: false,
        y: false
      }
      const offset = {
        x: this.start.x - (this.origin.x - e.clientX),
        y: this.start.y - (this.origin.y - e.clientY)
      }
      if (this.direction === 'horizontal') {
        if (this.repeat === true) {
          inbounds.x = true
        } else {
          inbounds.x = offset.x < 0 && offset.x * -1 < this.bounds.w
        }
        if (this.movecontinue && inbounds.x) {
          this.start.x = offset.x
          this.start.y = 0
        }
      } else {
        if (this.repeat === true) {
          inbounds.y = true
        } else {
          inbounds.y = offset.y < 0 && offset.y * -1 < this.bounds.h
        }
        if (this.movecontinue && inbounds.y) {
          this.start.y = offset.y
          this.start.x = 0
        }
      }
      this.change(this.start.x, this.start.y)
      this.origin.x = e.clientX
      this.origin.y = e.clientY

      e.stopPropagation()
      return false
    },
    handle (e) {
      this.movecontinue = false
      this.$el.omousemove = null

      if (e.type === 'mousedown') {
        this.origin.x = e.clientX
        this.origin.y = e.clientY
        this.movecontinue = true
        this.$el.onmousemove = this.move
      }
      e.stopPropagation()
      return false
    },
    reset () {
      this.change(0, 0)
    },
    change (x, y) {
      this.start = {
        x,
        y
      }
      this.$el.style.backgroundPosition = `${x}px ${y}px`
      this.$emit('viewchange', this.getView())
    },
    setRotation (angle) {
      const x = this.width * angle / 360.0
      this.change(x, 0)
    },
    getView () {
      const angle = this.start.x * 360.0 / this.width
      return {
        x: this.start.x,
        y: this.start.y,
        angle,
        width: this.width,
        height: this.height,
        imageWidth: this.imageWidth,
        imageHeight: this.imageHeight
      }
    }
  }
}
</script>
