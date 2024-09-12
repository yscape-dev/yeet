package com.example;

import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.infra.Blackhole;

import java.util.List;
import java.util.concurrent.TimeUnit;
import java.util.stream.IntStream;

import static java.util.stream.Collectors.toList;

@State(Scope.Thread)
@BenchmarkMode(Mode.AverageTime)
@OutputTimeUnit(TimeUnit.MICROSECONDS)
@Warmup(iterations = 3, time = 100, timeUnit = TimeUnit.MILLISECONDS)
@Measurement(iterations = 5, time = 100, timeUnit = TimeUnit.MILLISECONDS)
@Fork(1)
public class SimpleBenchmark {
    List<Integer> list = IntStream.range(0, 1_000_000)
            .boxed()
            .collect(toList());

    @Benchmark
    public void inverse_sqrt(Blackhole blackhole) {
        for (int i = 0; i < list.size(); i++) {
            blackhole.consume(1.0 / Math.sqrt(list.get(i)));
        }
    }

    @Benchmark
    public void fast_inverse_sqrt(Blackhole blackhole) {
        for (int i = 0; i < list.size(); i++) {
            blackhole.consume(1.0 / fastInverseSqrt(list.get(i)));
        }
    }

    @Benchmark
    public void fast_inverse_sqrt2(Blackhole blackhole) {
        for (int i = 0; i < list.size(); i++) {
            blackhole.consume(1.0 / fastInverseSqrt2(list.get(i)));
        }
    }

    private static float fastInverseSqrt(float x) {
        float xhalf = 0.5f * x;
        int i = Float.floatToIntBits(x);
        i = 0x5f3759df - (i >> 1);
        x = Float.intBitsToFloat(i);
        x = x * (1.5f - xhalf * x * x);
        return x;
    }

    private static float fastInverseSqrt2(float x) {
        int i = Float.floatToIntBits(x);
        i = 0x5f3759df - (i >> 1);
        x = Float.intBitsToFloat(i);
        return x * (1.5f - 0.5f * x * x);
    }
}
